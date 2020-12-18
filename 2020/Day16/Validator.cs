using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.IO;
using System.Linq;
using System.Text.RegularExpressions;
using AoC.Tokeniser;

namespace AoC
{
    public class Validator
    {
        public List<TicketToken> Invoke(List<TicketToken> passengerTickets, List<RuleToken> rules)
        {
            var ticketScanningErrorRate = 0;
            var invalidTickets = new List<int>();

            foreach(var ticket in passengerTickets)
                foreach(var field in ticket.Fields)
                    if( ! PassesAllRules(field) )
                    {
                        AddToInvalidTickets(ticket);
                        ticketScanningErrorRate += field;
                    }


            Console.WriteLine($"\nticket scanning error rate: {ticketScanningErrorRate}\n");
            return passengerTickets.Where(t => ! invalidTickets.Contains(t.TicketId) ).ToList();


            bool PassesAllRules(int field)
            {
                foreach(var rule in rules)
                    if(PassesRule(field, rule))
                        return true;

                // Failed all rules
                return false;
            }

            bool PassesRule(int field, RuleToken rule) =>
                   ( field >= rule.FirstRangeLowerBound  && field <= rule.FirstRangeUpperBound  )
                || ( field >= rule.SecondRangeLowerBound && field <= rule.SecondRangeUpperBound )
            ;

            void AddToInvalidTickets(TicketToken validTicket)
            {
                if( ! invalidTickets.Contains(validTicket.TicketId) )
                    invalidTickets.Add(validTicket.TicketId);
            }
        }
    }
}
