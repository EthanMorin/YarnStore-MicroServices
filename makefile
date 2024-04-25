up:
	sudo docker-compose up -d --build --remove-orphans 

upterm:
	sudo docker-compose up --build --remove-orphans

down:
	sudo docker-compose down





  # mongo:
  #   image: mongo:latest
  #   command: ["--replSet", "rs0", "--bind_ip_all", "--port", "27017"]
  #   ports:
  #     - 27017:27017
  #   extra_hosts:
  #     - "host.docker.internal:host-gateway"
  #   healthcheck:
  #     test: test $$(echo "rs.initiate({_id:'my-replica-set',members:[{_id:0,host:\"mongo:27017\"},{_id:1,host:\"mongo1:27018\"},{_id:2,host:\"mongo2:27019\"}]}).ok || rs.status().ok" | mongo --port 27017 --quiet) -eq 1
  #     interval: 10s
  #     start_period: 30s
  #   volumes:
  #     - "mongodb_data:/data/db"
  #   networks:
  #     - yarn_exchange

  # mongo1:
  #   image: mongo:latest
  #   command: ["--replSet", "rs0", "--bind_ip_all", "--port", "27018"]
  #   ports:
  #     - 27018:27018
  #   extra_hosts:
  #     - "host.docker.internal:host-gateway"
  #   volumes:
  #     - "mongodb1_data:/data/db"
  #   networks:
  #     - yarn_exchange

  # mongo2:
  #   image: mongo:latest
  #   command: ["--replSet", "rs0", "--bind_ip_all", "--port", "27019"]
  #   ports:
  #     - 27019:27019
  #   extra_hosts:
  #     - "host.docker.internal:host-gateway"
  #   volumes:
  #     - "mongodb2_data:/data/db"
  #   networks:
  #     - yarn_exchange