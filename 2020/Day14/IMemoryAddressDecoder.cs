using System;
using System.Collections.Generic;
using System.Linq;
using System.IO;


namespace AoC
{
    public interface IMemoryAddressDecoder
    {
        public IEnumerable<long> Decode(int address, string mask);
    }
}
