using System;
using System.Collections.Generic;
using System.Linq;
using System.IO;


namespace AoC
{
    public interface IConverter
    {
        string Convert(string from, string mask, long applyValue);
    }
}
