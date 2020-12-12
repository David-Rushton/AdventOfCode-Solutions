using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.Linq;


namespace AoC
{
    public enum Orintation
    {
        North = 0,
        East = 90,
        South = 180,
        West = 270
    }

    public record FerryLocation(
        Orintation Orintation,
        int NorthSouth,
        int EastWest
    );

    public record WaypointLocation(
        int NorthSouth,
        int EastWest
    );


    public interface IFerry
    {
        public FerryLocation PlotCourse(IEnumerable<NavigationInstruction> Instructions);
    }
}
