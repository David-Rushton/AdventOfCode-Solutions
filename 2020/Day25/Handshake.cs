using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.Linq;


namespace AoC
{
    public class Handshake
    {
        public void Invoke(int subjectNumber, long cardPublicKey, long doorPublicKey)
        {
            var cardLoopSize = GetLoopSize(7, cardPublicKey);
            var doorLoopSize = GetLoopSize(7, doorPublicKey);
            var encryptionKey = GetEncryptionKey(cardLoopSize, doorPublicKey);

            Console.WriteLine
            (
                string.Format
                (
                    "\nResult\n  Card loop size: {0}\n  Door loop size: {1}\n  Encryption key: {2}",
                    cardLoopSize,
                    doorLoopSize,
                    encryptionKey
                )
            );
        }


        private long GetLoopSize(long subjectNumber, long publicKey)
        {
            long value = 1;
            var loopCounter = 0;

            while(value != publicKey)
            {
                value = value * subjectNumber;
                value = value % 20201227;
                loopCounter++;
            }

            return loopCounter;
        }

        private long GetEncryptionKey(long loopSize, long subjectNumber)
        {
            long value = 1;

            for(var loop = 0; loop < loopSize; loop++)
            {
                value = value * subjectNumber;
                value = value % 20201227;
            }

            return value;
        }
    }
}
