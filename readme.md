Ce projet est un test technique pour un poste de développeur backend Golang en alternance.
Il s'agit d'une API REST qui permet de faire un CRUD de document.

## Dependances
- Golang 1.23
- Docker
- GitHub actions

## Installation
Pour installer le projet, il suffit d'installer docker sur votre machine et de lancer la commande suivante:
```docker pull aives81/vade-doc-mngt-api:latest```

## Utilisation
Pour lancer le projet, il suffit de lancer la commande suivante:
```docker run -p 8001:8001 aives81/vade-doc-mngt-api:latest```

# Routes
### Création d'un document
- POST /documents
#### Le body de la requête doit être de la forme suivante:
```json
{
  "id": "identifiant",
  "name": "nom du document",
  "description": "description du document",
}
```
### Liste des documents
- GET /documents
### Détails d'un document
- GET /documents/{id}