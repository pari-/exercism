# exercism

Convenience layer for my exercism integration

<!-- toc -->

- [Description](#description)
- [Dependencies](#dependencies)
- [License](#license)
- [Author Information](#author-information)

<!-- tocstop -->

## Description

Since I do not want to publically expose my API key, but still benefit for using
a public git repostory, used as main storage layer for exercism configuration(s)
as well as exercism track progress I use this convenience layer for integration.

What it does:
- fetch upstream exercism\_completion.bash in case it's missing
- enable bash completion for exercism
- wrap gpg-encrypted configuration decryption at exercism-cli runtime

## Integration

I use the following snippet in my `~/.bash_profile`:

```bash
#
# exercism
#
EXERCISM_DIR="/dir/of/your/choice"
if [ -f "${EXERCISM_DIR}/config/bash_profile.sh" ]; then
  source "${EXERCISM_DIR}/config/bash_profile.sh" "${EXERCISM_DIR}"
fi
```

## License

MIT

## Author Information

* Patrick Ringl
