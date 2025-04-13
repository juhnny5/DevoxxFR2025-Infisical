# Démonstrations

La plateforme de démonstration est disponible [ici](https://infisical.devoxx.jbriault.fr/).

<iframe src="https://giphy.com/embed/ZtAIbzohYzyblovfwV" width="480" height="480" style="" frameBorder="0" class="giphy-embed" allowFullScreen></iframe><p><a href="https://giphy.com/stickers/hammer-wreck-pistolavistas-ZtAIbzohYzyblovfwV"></a></p>

## Pour bien commencer

```bash
# Connecter la CLI à Infisical
brew update && brew install infisical

# Se connecter
infisical login --domain https://infisical.devoxx.jbriault.fr

# En alternative (via une variable d'env)
export INFISICAL_API_URL="https://infisical.devoxx.jbriault.fr"
infisical login

# Les informations sont stockées ici
cat ~/.infisical/infisical-config.json
```

## Explications

### Initialiser un nouveau projet

```bash
# Dans un projet existant
infisical init

# Créé localement un fichier qui ne contient pas d'informations sensibles
.infisical.json
```

### Récupérer un secret en fonction de son nom et de l'environnement

```bash
# Etape 1 : Générer token pour un environnement donné sur son projet
# Etape 2 : Lancer le curl ;)
curl -sr GET \
  --url https://infisical.devoxx.jbriault.fr/api/v3/secrets/raw/<NOM_DU_SECRET> \
  --header 'Authorization: Bearer <TOKEN>' | jq .secret
```

### Afficher les secrets via la CLI

```bash
cd demo/project_example

# On peut générer un fichier .env avec ses secrets en variable d'env
infisical secrets generate-example-env --env dev > .env

# Récupérer un secret (par défaut sur l'env de dev)
infisical secrets get API_KEY

# Récupérer un secret sur l'env de prod (prd)
infisical secrets get API_KEY --env prd
```

### Injecter des secrets via la CLI

```bash
cd demo/project_example/

infisical secrets set API_KEY=devsjdgwkeudyjwe --env dev
infisical secrets set API_KEY=prdsjdgwkeudyjwe --env prd
```

### Récupérer les secrets dans un projet Go

```bash
cd demo/go_get_secret/

# Charger les variables d'env
source .env

# Sinon vous pouvez le faire comme suit
echo "
export CLIENT_ID=''
export CLIENT_SECRET=''
export PROJECT_ID=''
export SITE_URL='https://infisical.devoxx.jbriault.fr'
" > .env

go run main.go
```

### Scanner un projet qui possède des secrets

```bash
cd demo/project_with_secrets/
# Scanner en mode verbeux
infisical scan -v

# Scanner les fichiers non commités
infisical scan git-changes

# Installer le hook pre-commit
infisical scan install --pre-commit-hook

# Ignorer des secrets
touch .infisicalignore
```
