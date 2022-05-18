package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// remake this struct
type (
	// modelo de BD para categorias de productos
	Categorias struct {
		ID primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
		// Nombre de la categoria
		CategoriaNombre string `json:"category_name" bson:"category_name"`
		// El titulo de la categoria en el HTML osea <title>...</titl>
		CategoriaTitulo string `json:"category_title" bson:"category_title"`
		// descripcion presente en HTML de la categoria osea el <meta name="description" content="..."
		CategoriaDescripcion string `json:"category_description" bson:"category_description"`
		// puede ser una carta o una tabla
		CategoriaTipo string `json:"category_type" bson:"category_type"`
	}

	// modelo de BD para sub categorias como en el caso del producto ultravioleta en la pagina: https://waterpurificationsupplies.com/uv-sterilizers/
	SubCategorias struct {
		ID primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
		// Nombre de la sub categoria
		CategoriaNombre string `json:"category_name" bson:"category_name"`
		// Descripcion de la sub categoria explicando de que trata esta sub categoria
		CategoriaDescripcion string `json:"category_description" bson:"category_description"`
		// puede ser una carta o una tabla
		CategoriaTipo string `json:"category_type" bson:"category_type"`
	}

	// modelo de BD para los productos en forma de tabla como en el caso de: https://waterpurificationsupplies.com/membranes
	Table struct {
		ID primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
		// el nombre del producto
		NombreProducto string `bson:"nombre_del_producto" json:"nombre_del_producto"`
		// descripcion del producto/tabla
		DescripcionTabla string `bson:"descripcion_tabla" json:"descripcion_tabla"`
		// imagen asociada con el producto
		ImagenTabla string `bson:"imagen_tabla" json:"imagen_tabla"`
		// informacion dentro de la tabla
		TableData []ColumnaDeTabla `bson:"table_data" json:"table_data"`
		// categoria de la tabla y sub categoria
		TableCategory    string `bson:"table_category" json:"table_category"`
		TableSubCategory string `bson:"table_sub_category" json:"table_sub_category"`
	}
	// estructura para almacenar los datos de la tabla
	ColumnaDeTabla struct {
		ID primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
		// header/titulo de la columna de la tabla que va explicar que pasa en esa columna
		TituloColumna string `bson:"titulo_columna" json:"titulo_columna"`
		// informacion de la columna/table data
		InfoColumna string `bson:"inf_columna" json:"inf_columna"`
	}

	Carta struct {
		ID primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
		// Nombre del producto para poner en la carta
		CartaTitulo string `bson:"carta_titulo" json:"carta_titulo"`
		// Descripcion del producto
		CartaDesc string `bson:"carta_desc" json:"carta_desc"`
		// Imagen para asociar con la carta
		CartaImagen string `bson:"carta_imagen" json:"carta_imagen"`
		// (opcional) Un link a un pdf o doc para ver mas info sobre el producto en cuestion
		MasInfoCarta string `bson:"mas_info_carta" json:"mas_info_carta"`
		// Categoria y sub categoria de la carta
		CategoriaCarta    string `bson:"categoria_carta" json:"categoria_carta"`
		SubCategoriaCarta string `bson:"sub_categoria_carta" json:"sub_categoria_carta"`
	}
)

func ConectarBDCategorias() (*mongo.Collection, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return client.Database("wps-db").Collection("categorias"), nil
}
func ConectarBDTablas() (*mongo.Collection, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return client.Database("wps-db").Collection("tablas"), nil
}
func ConectarBDCartas() (*mongo.Collection, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return client.Database("wps-db").Collection("cartas"), nil
}
