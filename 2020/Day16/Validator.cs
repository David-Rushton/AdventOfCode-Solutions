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
        public List<TicketToken> Invoke(TicketToken myTicket, List<TicketToken> passengerTickets, List<RuleToken> rules)
        {
            var ticketScanningErrorRate = 0;
            var validTickets = new List<TicketToken>();

            foreach(var ticket in passengerTickets)
                foreach(var field in ticket.Fields)
                    if(PassesAllRules(field))
                       validTickets.Add(ticket);
                    else
                       ticketScanningErrorRate += field;


            Console.WriteLine($"\nticket scanning error rate: {ticketScanningErrorRate}\n");
            return validTickets;


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
        }
    }
}
