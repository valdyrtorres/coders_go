Mongo
go get gopkg.in/mgo.v2

go run main.go

Iniciar o mongodb

mongod –-dbpath=C:\data\db 

ex: mongod --dbpath=C:\valdir\private\coders_go\restgo\data\db 

Comandos uteis

db.help()
Lista todos os comandos que podem ser aplicados no banco de dados e contém também  uma explicação (guarde este comando no coração);

show dbs
Mostra todos os bancos de dados existentes no servidor.

db
Mostra qual banco de dados está sendo usado. Se você acabou de instalar o MongoDB, e executar este comando verá que o banco de dados usado é o test, banco de dados default e para testes do MongoDB

use nomeBD
Este comando acessa um determinado banco de dados, quando ele existe. Se o banco não existe ele é criado automaticamente.

mongo
é o cliente

show collections

db.customers.find()

db.customers.insert({ nome: "Luiz", idade: 29 })

GOPATH:
C:\valdir\private\coders_go\go\src\projetos\go-restapi>set gopath
GOPATH=C:\valdir\private\coders_go\go

TESTANDO:
go run main.go no seu terminal, em seguida abra o endereço http://localhost:3000/api/v1/movies no seu navegador.

1) Criar (POST): OK
http://localhost:3000/api/v1/movies
{
  "name": "The Equalizer",
  "description": "Robert McCall (Denzel Washington), a man of mysterious origin who believes he has put the past behind him, dedicates himself to creating a quiet new life. However, when he meets Teri (Chloë Grace Moretz), a teenager who has been manhandled by violent Russian mobsters, he simply cannot walk away. With his set of formidable skills, McCall comes out of self-imposed retirement and emerges as an avenging angel, ready to take down anyone who brutalizes the helpless.",
  "thumb_image": "http://t2.gstatic.com/images?q=tbn:ANd9GcQkGfxoavBEp4fR6P-yi2mIkUl1aZHHFIietLK4GriI5YyvGSJ7",
  "active": true
}

2) GetAll (busca todos os registros) (get): OK
http://localhost:3000/api/v1/movies

3) GetById OK
http://localhost:3000/api/v1/movies/"id"
Ex: http://localhost:3000/api/v1/movies/5d76578987cb4f4decacb8cd

4) Update (PUT) OK
http://localhost:3000/api/v1/movies/5d7694eb87cb4f7cf0a0adb9
{
  "id":"5d7694eb87cb4f7cf0a0adb9",
  "name": "The Equalizer 2",
  "description": "If you have a problem and there is nowhere else to turn, the mysterious and elusive Robert McCall will deliver the vigilante justice you seek. This time, however, McCall's past cuts especially close to home when thugs kill Susan Plummer -- his best friend and former colleague. Now out for revenge, McCall must take on a crew of highly trained assassins who'll stop at nothing to destroy him.",
  "thumb_image": "http://t1.gstatic.com/images?q=tbn:ANd9GcRCss5jFvU87fla4jEIwpv9dUAdZKzeUYHY_mhqzDxrvX9ppWyJ",
  "active": true
}

5) DELETE OK
http://localhost:3000/api/v1/movies/"id"
Ex: http://localhost:3000/api/v1/movies/5d76578987cb4f4decacb8cd

TESTES COM O CURL:
1)Consulta: OK
curl -v -X GET http://localhost:3000/api/v1/movies    

2)Criar: OK
curl -v POST -d "{\"name\":\"Teste curl dois\", \"description\":\"descricao curl dois\", \"thumb_image\": \"http://t2.gstatic.com/images?q=tbn:ANd9GcQkGfxoavBEp4fR6P-yi2mIkUl1aZHHFIietLK4GriI5YyvGSJ7\", \"active\": true }" http://localhost:3000/api/v1/movies

3)Update OK
curl -v -X PUT -d "{\"id\":\"5d76a44887cb4f8b5cc4135e\",\"name\":\"Rico curl Torres\", \"description\":\"descricao update curl dois\", \"thumb_image\": \"http://t2.gstatic.com/images?q=tbn:ANd9GcQkGfxoavBEp4fR6P-yi2mIkUl1aZHHFIietLK4GriI5YyvGSJ7\", \"active\": true }" http://localhost:3000/api/v1/movies/5d76a44887cb4f8b5cc4135e

Delete
curl -v -X DELETE http://127.0.0.1:8701/tasks/5d6e9c1487cb4f2870c1b7d9



