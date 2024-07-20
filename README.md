# Getting Started with Buf gRPC Gateway

Follow the instructions in the [Buf documentation](https://buf.build/docs/tutorials/getting-started-with-buf-cli#fix-lint-failures) to install Buf and understand the basics of the Buf CLI.


## Install Buf and Protoc

```bash
brew install protoc
brew install bufbuild/buf/buf
```

## Install plugins

```bash
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

## Generate the gRPC Gateway

```bash
cd $HOME/getting-started-with-buf-cli
```

```bash
make generate
```

> We can run `buf lint` and `buf breaking` to check for linting errors and breaking changes.



## Run the gRPC Gateway

```bash
make run
```

```bash
gRPC call using buf
```bash
buf curl \
--schema . \
--data '{"pet_type": "PET_TYPE_SNAKE", "name": "Ekans"}' \
http://localhost:2080/pet.v1.PetStoreService/PutPet

{
  "pet": {
  "petType": "PET_TYPE_SNAKE",
  "petId": "11f103d3-4d28-451b-8a50-db67f2b50302",
  "name": "Ekans"
  }
}
```

Http call using httpie
```bash
http PUT http://localhost:9085/v1/pets petType="PET_TYPE_SNAKE" name="Ekans"

{
    "pet": {
        "createdAt": null,
        "name": "Dora",
        "petId": "295199fa-d027-40b3-b154-d888122d91ad",
        "petType": "PET_TYPE_SNAKE"
    }
}
```