package log_rus

import (
	"fmt"
	"os"
	"strings"

	"github.com/fabiosebastiano/golang-microservices/src/api/config"
	"github.com/sirupsen/logrus"
)

/*
* Implementata come INTERFACCIA pubblica che espone i metodi di scrittura dei log SENZA riferimento diretto alla libreria usata
 */
var (
	//SI PUO' MODIFICARE LA LIBRERIA E NON DOBBIAMO PREOCCUPARCI DI AGGIORNARE IL CODICE IN GIRO!
	Log *logrus.Logger
)

func init() {
	level, err := logrus.ParseLevel(config.Loglevel)

	if err != nil {
		level = logrus.DebugLevel
	}
	//3 CONFIGURAZIONI DI BASE:
	Log = &logrus.Logger{
		Level:     level,                   //1) LIVELLO
		Out:       os.Stdout,               //2) OUTPUT
		Formatter: &logrus.JSONFormatter{}, //3) TIPO FORMATTER --> si può differenziare per tipo di ambiente
	}

	/* DIFFERENZIAZIONE FORMATTER IN FUNZIONE DELL'AMBIENTE
	if config.IsProduction() {
		Log.Formatter = &logrus.JSONFormatter{}
	} else {
		Log.Formatter = &logrus.TextFormatter{}
	}
	*/
}

func Info(msg string, tags ...string) {
	if Log.Level < logrus.InfoLevel {
		return
	}
	//parsifichiamo la lista di tag in WithFields e ci aggiungiamo il messaggio
	Log.WithFields(parseFields(tags...)).Info(msg)
}

//Error > farsi passare anche errore per formattarlo e poi loggarlo
func Error(msg string, err error, tags ...string) {
	if Log.Level < logrus.ErrorLevel {
		return
	}

	msg = fmt.Sprintf("%s - ERROR - %v", msg, err)

	//parsifichiamo la lista di tag in WithFields e ci aggiungiamo il messaggio
	Log.WithFields(parseFields(tags...)).Error(msg)
}
func Debug(msg string, tags ...string) {
	if Log.Level < logrus.DebugLevel {
		return
	}
	//parsifichiamo la lista di tag in WithFields e ci aggiungiamo il messaggio
	Log.WithFields(parseFields(tags...)).Debug(msg)
}

func parseFields(tags ...string) logrus.Fields {

	result := make(logrus.Fields, len(tags))
	for _, tag := range tags {
		els := strings.Split(tag, ":")
		result[strings.TrimSpace(els[0])] = strings.TrimSpace(els[1])
	}

	return result
}
