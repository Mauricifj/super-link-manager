# Super Link Manager

More information about Super Link Cielo [here](https://desenvolvedores.cielo.com.br/api-portal/pt-br/content/super-link-cielo).

### Prerequisites

- Docker

### GitHub clone

Clone this repository
```shell script
git clone https://github.com/Mauricifj/super-link-manager
```

### Credentials

Set your credentials on _.env_ file

```dotenv
USERNAME=YOUR-USERNAME-GOES-HERE
PASSWORD=YOUR-PASSWORD-GOES-HERE
...
```

### Build and run

Run the following command on your terminal from the cloned repository

```shell script
docker-compose up --build -d
```

### Usage

Go to [localhost:8000](http://localhost:8000) on your browser to access Super Link Manager. 