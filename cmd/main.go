package main

import (
	"github.com/amrchnk/account_service/pkg/handler"
	"github.com/amrchnk/account_service/pkg/repository"
	"github.com/amrchnk/account_service/pkg/service"
	pb "github.com/amrchnk/account_service/proto"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

func init(){
	err := initConfig()
	if err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	err = godotenv.Load(filepath.Join("./", ".env"))
	if err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}
}

func main(){
	con, err := net.Listen("tcp", ":"+viper.GetString("port"))
	if err != nil {
		log.Fatalf("tcp connection error: %v", err)
	}

	grpcServer := grpc.NewServer()

	repo := repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("PASSWORD"),
	}

	db, err := repository.NewPostgresDb(repo)
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)

	pb.RegisterAccountServiceServer(grpcServer, handler.NewService(services))
	go func() {
		if err := grpcServer.Serve(con); err != nil {
			log.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()
	log.Printf("Auth service started at port :%s", viper.GetString("port"))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := grpcServer.Stop; err != nil {
		log.Fatal("Server has been stopped")
	}
	if err := db.Close(); err != nil {
		log.Fatalf("error occured on db connection close: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}