# SaaS de Videomonitoramento Inteligente 

Sistema multiplataforma de monitoramento de segurança que utiliza inteligência artificial para detectar intrusões ou comportamentos suspeitos via câmeras, acionando notificações em tempo real para as autoridades ou responsáveis. Nesse caso sera com camera instaladas em HUBS em breve para monitoramento em outras tecnlogias 

---

## 🧠 Arquitetura Geral

[Câmera IP / Webcam] → [Detector Python (YOLO + OpenCV)] → gRPC → [Backend Go (Core + API)] → front end (em construcao)
                                                                      |
                                                            [PostgreSQL + em construcao com o que mais sera colocado]

- **Detector Python**: captura vídeo, usa YOLOv8 para detectar pessoas, envia alertas via gRPC.
- **Backend Go**: gerencia usuários, câmeras, eventos e notificações. Expõe API gRPC.
- parte visual do(UI) ainda sera pesado qual a melhor forma.


## 🛠️ Tecnologias

| Componente | Tecnologia |
|------------|------------|
| Backend | Go 1.21+ |
| Comunicação | gRPC (protobuf) |
| IA / Visão | Python 3.12, OpenCV, Ultralytics YOLOv8 |
| Banco de dados | PostgreSQL (em construcao) |
| Notificações | Firebase Cloud Messaging (em construcao , cogitando) |
| Frontend (futuro) | Flutter(cogitando) (Android, iOS, Desktop) ideal que nativamente funcione em todas as platafotrmas com a mesma eficiencia |


## 📁 Estrutura do Projeto

saas-monitoramento/
├── backend/              # Servidor Go
│   ├── cmd/server/main.go
│   ├── internal/
│   │   ├── alert/service.go
│   │   └── notify/notify.go
│   ├── proto/alert.proto
│   ├── gen/              # Código Go gerado
│   └── go.mod
├── detector/             # Módulo de IA (Python)
│   ├── detector.py
│   ├── proto/alert.proto
│   ├── gen/              # Stubs Python gerados
│   ├── requirements.txt
│   └── .venv/            # Ambiente virtual
├── .gitignore
└── README.md

## 🚀 Como Executar

### Pré-requisitos
- Go 1.21+
- Python 3.12+ com venv e pip
- Git

### 1. Clonar e iniciar o Backend Go

cd backend
go mod tidy
go run ./cmd/server

O servidor rodará em localhost:50051.

### 2. Rodar o Detector Python

cd detector
python3 -m venv .venv
source .venv/bin/activate
pip install -r requirements.txt

Ajuste a câmera no detector.py (webcam: RTSP_URL = 0("seria para exemplo que foi usado para camera da webcam do notebook")) e execute:

python detector.py

### 3. Ver o fluxo

Com ambos rodando, ao detectar uma pessoa:
- Terminal Go: [ALERTA RECEBIDO] ...  
- Terminal Python: ALERTA ENVIADO! ID: ...

Pressione 'q' na janela de vídeo para sair.



## 📈 Próximos Passos
a combinar

## 📜 Licença
MIT

---

Desenvolvido como prova de conceito de SaaS de monitoramento inteligente.

