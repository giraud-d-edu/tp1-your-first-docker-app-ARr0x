# Utiliser l'image de base officielle de Go
FROM golang:1.21.3-alpine

# Installer les dépendances nécessaires pour SQLite
RUN apk add --no-cache gcc musl-dev

# Définir le répertoire de travail dans le conteneur
WORKDIR /app

# Copier les fichiers go.mod et go.sum dans le répertoire de travail
COPY go.mod go.sum ./

# Télécharger les dépendances
RUN go mod download

# Copier le reste du code source dans le répertoire de travail
COPY . .

# Construire l'application Go avec CGO activé
RUN CGO_ENABLED=1 go build -o main .

# Exposer le port sur lequel l'application va tourner
EXPOSE 8080

# Commande pour exécuter l'application
CMD ["./main"]
