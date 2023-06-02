package database

import (
	"context"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func Connect() {
	dbHost := os.Getenv("DB_HOST")
	// dbPort :=  os.Getenv("DB_PORT")
	// TODO เอาไปแปลงเป็นstring ใน setting
	dbPort := "27017"
	// dbName := os.Getenv("DB_NAME")
	// dbUsername := os.Getenv("DB_USERNAME")
	// dbPassword := os.Getenv("DB_PASSWORD")

	// 1x เก็บ ตั้งค่า ข้อมูล พวกhost port ที่จะเชื่อมกับuri mongo
	// 1.1	mongo จะมีdefult ที่ให้เชื่อม คือ mongo://host:port
	connectionURI := fmt.Sprintf("mongodb://%s:%s", dbHost, dbPort)
	// 1.2 เป็น reciver method คือก้อนobject ที่การกำหนดค่าต่างๆตามที่เรากำหนด options.Client()การกำหนดค่าต่างๆของไคลเอ็นต์
	// 1.3 ApplyURI ฟังก์ที่ใช้ในการกำหนด uri ของฐานข้อมูลที่จะเชื่อมต่อ
	clientOptions := options.Client().ApplyURI(connectionURI)
	// 2x เชื่อมต่อกับmongo ใช้pkg mongo
	// 2.1 context ทำให้ คุมการทำงานของprocessนั้นได้ กรณีนี้  ใช้contexแค่ตรวจสอบสถานะของการเชื่อมต่อ

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		logrus.Fatal(err)
	}
	// สร้างหรือเรียกDB
	DB = client.Database("golang-test")
}
