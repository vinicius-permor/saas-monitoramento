// Package alert para envio de alerta para o usuario
package alert

import (
	"context"
	"log"
	pb "saas-monitoramento/backend/gen"
	"saas-monitoramento/backend/internal/notify"

	"github.com/google/uuid"
)

// AlertServer implementa o servidor grpc gerado
type AlertServer struct {
	pb.UnimplementedAlertServiceServer
}

func (s *AlertServer) SendAlert(ctx context.Context, req *pb.SendAlertRequest) (*pb.SendAlertResponse, error) {
	// Implementação do método SendAlert
	// o alerta vai trazer camera, tipo , a hora da ocorrencia

	alertID := uuid.New().String()
	log.Printf("[ALERTA RECEBIDO] Camera:  %s , tipo: %s , Hora: %v", req.CameraId, req.ThreatType, req.Timestamp)
	notify.DispatchAlert(alertID, req.CameraId, req.ThreatType, req.SnapshotUrl, req.Timestamp)
	return &pb.SendAlertResponse{
		AlertId: alertID,
	}, nil
}

func (s *AlertServer) StreamAlerts(req *pb.StreamAlertsRequest, stream pb.AlertService_StreamAlertsServer) error {
	for {
		select {
		case <-stream.Context().Done():
			return nil
		}
	}
}
