#include "openctm.h"

extern CTMuint readerHelper(void * aBuf, CTMuint aCount, void * aUserData);

void ctmLoadStream(CTMcontext aContext, void * aUserData) {
  ctmLoadCustom(aContext, &readerHelper, aUserData);
}
