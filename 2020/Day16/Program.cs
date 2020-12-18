using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.IO;
using System.Linq;
using System.Text.RegularExpressions;
using AoC.Tokeniser;

// Ticket Translation
// .--------------------------------------------------------.
// | ????: 101    ?????: 102   ??????????: 103     ???: 104 |
// |                                                        |
// | ??: 301  ??: 302             ???????: 303      ??????? |
// | ??: 401  ??: 402           ???? ????: 403    ????????? |
// '--------------------------------------------------------'


namespace AoC
{
    class Program
    {
        static void Main(string[] args)
        {
            var app = Bootstrap(args.Contains("--test"));
            var validTickets = app.Validator.Invoke
                (
                    app.InputReader.PassengerTickets,
                    app.InputReader.Rules
                )
            ;

            app.FieldDetector.Invoke
                (
                    app.InputReader.MyTicket,
                    validTickets,
                    app.InputReader.Rules
                )
            ;
        }


        private static (InputReader InputReader, Validator Validator, FieldDetector FieldDetector) Bootstrap(bool useTestInput) =>
            (
                new InputReader(useTestInput),
                new Validator(),
                new FieldDetector()
            )
        ;
    }
}
