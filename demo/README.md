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
```
