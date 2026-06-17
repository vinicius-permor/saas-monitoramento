// Package notify gera a notificacao para o usuario
package notify

import "log"

// DispatchAlert essa funcao sera responsavel para envio do alerta contendo
// contendo id do alerta , id da camra , tipo  , url para visualizacao , tempo e print da imagem
func DispatchAlert(alertID, cameraID, threatType, snapshotURL string, timestamp int64) {
	log.Printf("[ALERTA DISPARADO] ID: %s, Camera: %s, Tipo: %s,URL: %s", alertID, cameraID, threatType, snapshotURL)
}
