install:
  @echo "Installing readme generator..."
  @cp ./bin/readme /usr/local/bin/readme
  @sleep 2
  @echo "Instalation is done."

.PHONY: install