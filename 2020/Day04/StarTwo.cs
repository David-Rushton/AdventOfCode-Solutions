using System;
using System.Collections.Generic;
using System.Linq;
using System.IO;


namespace AoC
{
    public class StarTwo: IStar
    {
        public void Invoke(PassportReader reader) => Console.WriteLine(reader.ProcessPassports());
    }
}
