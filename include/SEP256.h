#include <stdbool.h>

extern bool GenerateSEP256Key(char *buffer, int *length);
extern bool SEP256KeyAgreement(char *privateKey, int privateKeyLength, char *publicKey, int publicKeyLength, char *buffer, int *length);
