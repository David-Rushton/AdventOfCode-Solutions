using System;
using System.Collections.Generic;
using System.Linq;
using System.IO;


namespace AoC
{
    public record RunResult(
        int Right,
        int Down,
        int TreesEncountered
    );


    public class StarTwo: IStar
    {
        public void Invoke(PassportReader reader)
        {
            throw new NotImplementedException();
        }
    }
}
