using System;
using System.Linq;


namespace AoC
{
    class Program
    {
        static void Main(string[] args)
        {
            var useTestKey = args.Contains("--test");
            var keys = GetPublicKeys(useTestKey);

            new Handshake().Invoke(7, keys.cardPublicLKey, keys.doorPublicKey);
        }


        private static (long cardPublicLKey, long doorPublicKey) GetPublicKeys(bool useTestKeys)
        {
            return
                useTestKeys
                ? (5764801, 17807724)
                : (5290733, 15231938)
            ;
        }
    }
}
