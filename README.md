# Challenge-golang-flutter

Projet challenge-golang-flutter

ESGI VOYAGE

Description
ESGI VOYAGE est une application conçue pour organiser des voyages en proposant diverses destinations. Elle permet aux utilisateurs de créer, sauvegarder et partager leurs voyages. Les utilisateurs peuvent également créer des groupes de voyage pour partager des informations et inviter d'autres personnes à rejoindre ces groupes.

Fonctionnalités
Créer un voyage : Les utilisateurs peuvent créer un nouveau voyage en ajoutant des détails tels que la destination, les dates, et d'autres informations pertinentes.
Sauvegarder les voyages : Tous les voyages créés sont sauvegardés pour permettre aux utilisateurs de les consulter plus tard.
Créer un groupe de voyage : Les utilisateurs peuvent créer des groupes pour chaque voyage afin de partager des informations et des ressources avec d'autres membres du groupe.
Partager des informations : Les membres d'un groupe de voyage peuvent partager des messages et des fichiers avec les autres membres du groupe.
Inviter des personnes : Les utilisateurs peuvent inviter d'autres personnes à rejoindre leur groupe de voyage en entrant leur adresse e-mail.


## Equipe de développement

| Nom         | Prenom   | Pseudo Git     |
|-------------|----------|----------------|
| Gomet       | Eddy     | EddyGmt        |
| Youssouf    | Nourdine | nono231000     |
| Traore      | Chiekhou | chiekhou       |


## Fonctionnalités et membres

| Fonctionnalité                                                                                                          | Développeur                                 |
|-------------------------------------------------------------------------------------------------------------------------|---------------------------------------------|
| Mise en place des apis utilitaires Auth(register , login ,reset-password,profil) et Groupe Voyage                       | Eddy Gomet (front et back)                  |
| Mise en place du chat et l'envoi de notifications push lors d'envoi de message et du la creation d'un groupe dans l'app | Eddy (back et front)                        |
| Mise en place de l'invitation par mail pour rejoindre un groupe                                                         | Eddy (back et front)                        | 
| Mise en place des api Destinations , Voyages , Activités ,                                                              | Chiekhou (front et back)                    |
| Mise en place du panel Admin                                                                                            | Chiekhou  (front et back)                   |                
| Mise en place du site mutilingue/Mise en place la fonctionnalité feature-flipping                                       | Chiekhou  (front et back)                   |
| Mise en place des tests                                                                                                 | Nourdine (back)                             | 

## Pour lancer le projet :

Pour le back se rendre dans le dossier /GO:

Pour installer la base de donnée on utilise un docker compose , executer la commande suivante : docker compose up -d

Pour lancer le serveur sur le localhost:8080:

1) go mod tiny (pour installer les dépendances)
2) go run server main.go
3) go run internal/migrate/migrate.go (pour executer des migrations)
4) swag init --parseDependency --parseInternal --parseDepth 1
5) go test / go test -v


Pour le front se rendre dans le dossier FLUTTER-APP:
flutter run pour lancer son projet selon l'environnement de son choix et d'avoir au préalable toute la coniguration pour lancer un émulateur android ou ios .




