# Démonstrations

La plateforme de démonstration est disponible [ici](https://infisical.devoxx.jbriault.fr/).

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

# Créé localement un fichier qui ne contient pas d'informations
.infisical.json
```

### Récupérer un secret en fonction de l'environnement

```bash
# Etape 1 : Générer token pour un environnement donné sur son projet
# Etape 2 : Lancer le curl ;)
curl -sr GET \
  --url https://infisical.devoxx.jbriault.fr/api/v3/secrets/raw/<NOM_DU_SECRET> \
  --header 'Authorization: Bearer <TOKEN>' | jq .secret
```
