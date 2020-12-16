using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.IO;
using System.Linq;
using System.Text.RegularExpressions;


namespace AoC.Tokeniser
{
    public class InputReader
    {
        const string _rulePattern = @"[a-z]+(\s[a-z]+)?:\s\d+-\d+\sor\s\d+-\d+";
        const string _ticketPattern = @"(\d+,){2,19}\d+";
        readonly string _testFilePath = Path.Join(Directory.GetCurrentDirectory(), "Test-Input.txt");
        readonly string _puzzleFilePath = Path.Join(Directory.GetCurrentDirectory(), "Input.txt");


        public InputReader(bool useTestData) => Initialise(useTestData);


        public List<RuleToken> Rules { get; private set; }

        public TicketToken MyTicket { get; private set; }

        public List<TicketToken> PassengerTickets { get; private set; }


        private void Initialise(bool useTestData)
        {
            var filePath = useTestData ? _testFilePath : _puzzleFilePath;
            var input = File.ReadAllText(filePath);

            InitialiseRules(input);
            InitialiseTickets(input);
        }

        private void InitialiseRules(string input) =>
            Rules = new Regex(_rulePattern).Matches(input)
                .Select
                (
                    r =>
                    {
                        var elements = r.Value.Split(':');
                        var fieldName = elements[0].Trim();
                        var numbers = elements[1].Split(" or ");
                        var firstNumbers = numbers[0].Trim().Split('-');
                        var secondNumbers = numbers[1].Trim().Split('-');

                        return new RuleToken
                        (
                                fieldName,
                                int.Parse(firstNumbers[0].Trim()),
                                int.Parse(firstNumbers[1].Trim()),
                                int.Parse(secondNumbers[0].Trim()),
                                int.Parse(secondNumbers[1].Trim())
                        );
                    }
                )
                .ToList()
            ;

        private void InitialiseTickets(string input)
        {
            var tickets = new Regex(_ticketPattern).Matches(input);

            // My ticket is always first
            MyTicket = new TicketToken(TicketOwner.Mine, GetFields(tickets[0].Value));


            // Rest belong to other passengers
            PassengerTickets = tickets
                .Skip(1)
                .Select( m => new TicketToken(TicketOwner.Passenger, GetFields(m.Value)) )
                .ToList()
            ;


            int[] GetFields(string fields) => fields.Split(',').Select(i => int.Parse(i)).ToArray();
        }
    }
}
