package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*TweetSeguidores es la estructura con la que devolveremos los tweets*/
type TweetSeguidores struct {
	ID                primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID            string             `bson:"usuarioid" json:"userId,omitempty"`
	UsuarioRelacionID string             `bson:"usuariorelacionid" json:"userRelationId"`
	//OJO Estra struct debe llamarse como la tabla
	Campos struct {
		ID      string    `bson:"_id" json:"_id,omitempty"`
		Mensaje string    `bson:"mensaje" json:"mensaje,omitempty"`
		Fecha   time.Time `bson:"fecha" json:"fecha,omitempty"`
	}
}
