Mongo
go get gopkg.in/mgo.v2

go run main.go

Iniciar o mongodb

mongod –-dbpath=C:\data\db 

ex: mongod --dbpath=C:\valdir\private\coders_go\restgo\data\db 

GOPATH:
C:\valdir\private\coders_go\go\src\projetos\go-restapi>set gopath
GOPATH=C:\valdir\private\coders_go\go

TESTANDO:
go run main.go no seu terminal, em seguida abra o endereço http://localhost:3000/api/v1/movies no seu navegador.

Criar (POST): OK
http://localhost:3000/api/v1/movies
{
  "name": "The Equalizer",
  "description": "Robert McCall (Denzel Washington), a man of mysterious origin who believes he has put the past behind him, dedicates himself to creating a quiet new life. However, when he meets Teri (Chloë Grace Moretz), a teenager who has been manhandled by violent Russian mobsters, he simply cannot walk away. With his set of formidable skills, McCall comes out of self-imposed retirement and emerges as an avenging angel, ready to take down anyone who brutalizes the helpless.",
  "thumb_image": "http://t2.gstatic.com/images?q=tbn:ANd9GcQkGfxoavBEp4fR6P-yi2mIkUl1aZHHFIietLK4GriI5YyvGSJ7",
  "active": true
}

GetAll (busca todos os registros) (get): OK
http://localhost:3000/api/v1/movies

GetById OK
http://localhost:3000/api/v1/movies/"id"
Ex: http://localhost:3000/api/v1/movies/5d76578987cb4f4decacb8cd

Update (PUT) OK
http://localhost:3000/api/v1/movies/5d7694eb87cb4f7cf0a0adb9
{
  "id":"5d7694eb87cb4f7cf0a0adb9",
  "name": "The Equalizer 2",
  "description": "If you have a problem and there is nowhere else to turn, the mysterious and elusive Robert McCall will deliver the vigilante justice you seek. This time, however, McCall's past cuts especially close to home when thugs kill Susan Plummer -- his best friend and former colleague. Now out for revenge, McCall must take on a crew of highly trained assassins who'll stop at nothing to destroy him.",
  "thumb_image": "http://t1.gstatic.com/images?q=tbn:ANd9GcRCss5jFvU87fla4jEIwpv9dUAdZKzeUYHY_mhqzDxrvX9ppWyJ",
  "active": true
}

DELETE OK
http://localhost:3000/api/v1/movies/"id"
Ex: http://localhost:3000/api/v1/movies/5d76578987cb4f4decacb8cd

