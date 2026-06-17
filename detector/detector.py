import sys
import os

sys.path.insert(0, os.path.join(os.path.dirname(__file__), "gen"))
import grpc
import cv2
import time
from ultralytics import YOLO
import alert_pb2
import alert_pb2_grpc

print("IMPORT OK")


# configucoes da camera
# esta sendo usado o exemplo com webcam do notebook
RTSP_URL = 0
GRPC_SERVER = "127.0.0.1:50051"
CAMERA_ID = "camera frontal 01"


# configuracao de carregamento do modelo YOLO ,
# responsalvel por detectar pessoas ,
# esse modele vem por padrao ja com treinamento de deteccao de pessoas.
model = YOLO("yolov8n.pt")


# configuracao de conexao com o servidor go
channel = grpc.insecure_channel(GRPC_SERVER)
alert_service_stub = alert_pb2_grpc.AlertServiceStub(channel)


def main() -> None:
    capture = cv2.VideoCapture(RTSP_URL)
    if not capture.isOpened():
        print("erro ao conectar na camera RTSP")
        return

    print("monitorando camera pressione 'q' para sair")

    intrusion_active = False
    cooldown_util = 0
    while True:
        ret, frame = capture.read()
        if not ret:
            print("err , frame perdido")
            continue

        # deteccao de objeto
        # usando classes[0] detecta somente pessoas
        # a medida do que precisar vai sendo escalonado "clasees[0] aumentando"
        results = model(frame, verbose=False, classes=[0])

        person_detected = False
        for result in results:
            if len(result.boxes) > 0:
                person_detected = True
                annotated = result.plot()
                cv2.imshow("Monitor", annotated)
                break
        if not person_detected:
            cv2.imshow("Monitor", frame)

        if person_detected:
            if not intrusion_active and time.time() > cooldown_util:
                intrusion_active = True
                # quantidade se segundos entre alertas
                cooldown_util = time.time() + 30

            # enviando alerta para o servidor em go

            request = alert_pb2.SendAlertRequest(
                camera_id=CAMERA_ID,
                threat_type="intrusion",
                snapshot_url="http://exemplo.com/foto.jpg",  # exemplo
                timestamp=int(time.time()),
            )

            response = alert_service_stub.SendAlert(request)
            print(f"ALERTS ENVIADO ID: {response.alert_id}")

        else:
            intrusion_active = False

        if cv2.waitKey(1) & 0xFF == ord("q"):
            break

    capture.release()
    cv2.destryAllWindows()


if __name__ == "__main__":
    main()
