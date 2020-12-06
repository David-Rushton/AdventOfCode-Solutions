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
            var validPasswordCount = 0;

            foreach(var password in input)
            {
                if(_passwordValidator.IsTobogganValidPassword(password))
                {
                    validPasswordCount ++;
                    Console.WriteLine($"Valid password found: {password}");
                }
            }

            Console.WriteLine($"\nFound {validPasswordCount} valid passwords");
        }
    }
}
