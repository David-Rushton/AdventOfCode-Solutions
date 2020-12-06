using System;
using System.Collections.Generic;
using System.Linq;
using System.IO;


namespace AoC
{
    public class StarTwo: IStar
    {
        readonly PasswordValidator _passwordValidator;


        public StarTwo(PasswordValidator passwordValidator) => (_passwordValidator) = (passwordValidator);


        public void Invoke(List<string> input)
        {
            throw new NotImplementedException();
        }
    }
}
