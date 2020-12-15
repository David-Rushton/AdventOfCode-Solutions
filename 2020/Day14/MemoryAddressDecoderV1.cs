using System;
using System.Collections.Generic;
using System.Linq;
using System.IO;


namespace AoC
{
    public class MemoryAddressDecoderV1: IMemoryAddressDecoder
    {
        public IEnumerable<long> Decode(int address, string mask)
        {
            yield return address;
        }
    }
}
