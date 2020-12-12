using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.Linq;


namespace AoC
{
    public class Ferry2: IFerry
    {
        public FerryLocation PlotCourse(IEnumerable<NavigationInstruction> Instructions)
        {
            var ferryLocation = new FerryLocation(Orintation.East, 0 , 0);
            var waypointLocation = new WaypointLocation(1, 10);

            foreach(var instruction in Instructions)
            {
                Debug.Assert(instruction.Direction != NavigationDirection.Unknown, $"Unsupported direction: {instruction}");
                Debug.Assert(instruction.Value >= 0, $"Negative instruction values are not supported: {instruction}");

                (waypointLocation, ferryLocation) = ProcessInstruction(waypointLocation, ferryLocation, instruction);
                ReportLocation(waypointLocation, ferryLocation, instruction);
            }


            return ferryLocation;
        }


        private (WaypointLocation Waypoint, FerryLocation Ferry) ProcessInstruction(
            WaypointLocation waypointLocation, FerryLocation ferryLocation, NavigationInstruction instruction
        )
        {
            if(instruction.Direction == NavigationDirection.Forward)
                return (waypointLocation, MoveFerry(waypointLocation, ferryLocation, instruction));

            return (MoveWaypoint(waypointLocation, instruction), ferryLocation);
        }

        private FerryLocation MoveFerry(
            WaypointLocation Waypoint, FerryLocation location, NavigationInstruction instruction
        ) =>
            new FerryLocation(
                location.Orintation,
                location.NorthSouth + (Waypoint.NorthSouth * instruction.Value),
                location.EastWest + (Waypoint.EastWest * instruction.Value)
            )
        ;

        private WaypointLocation MoveWaypoint(WaypointLocation location, NavigationInstruction instruction)
        {
            if(instruction.Direction is NavigationDirection.Left or NavigationDirection.Right)
                return RotateWaypoint();

            return OffsetWaypoint();


            // Move the waypoint.
            WaypointLocation OffsetWaypoint()
            {
                (int NorthSouth, int EastWest) offset = instruction.Direction switch
                {
                    NavigationDirection.North   => (instruction.Value,          0                       ),
                    NavigationDirection.East    => (0,                          instruction.Value       ),
                    NavigationDirection.South   => (instruction.Value * -1,     0                       ),
                    NavigationDirection.West    => (0,                          instruction.Value *-1   ),
                    _                           => (0,                          0                       )
                };

                return new WaypointLocation
                (
                    location.NorthSouth + offset.NorthSouth,
                    location.EastWest + offset.EastWest
                );
            }

            // Rotate it around the Ferry
            WaypointLocation RotateWaypoint()
            {
                (int NorthSouth, int EastWest) rotate = instruction.Direction switch
                {
                    NavigationDirection.Left   =>
                        instruction.Value switch
                        {
                            90  => (location.EastWest,          location.NorthSouth * -1  ),
                            180 => (location.NorthSouth * -1,   location.EastWest * -1    ),
                            270 => (location.EastWest * -1,     location.NorthSouth       ),
                            _   => (location.NorthSouth,        location.EastWest         )
                        },
                    NavigationDirection.Right    =>
                        instruction.Value switch
                        {
                            90  => (location.EastWest * -1,     location.NorthSouth       ),
                            180 => (location.NorthSouth * -1,   location.EastWest * -1    ),
                            270 => (location.EastWest,          location.NorthSouth * -1  ),
                            _   => (location.NorthSouth,        location.EastWest         )
                        },
                    _ => (0, 0)
                };

                return new WaypointLocation(rotate.NorthSouth, rotate.EastWest);
            }
        }

        private void ReportLocation(
            WaypointLocation waypointLocation, FerryLocation ferryLocation, NavigationInstruction instruction
        ) =>
            Verbose.WriteLine
            (
                String.Format
                (
                    "Course change:\n  Instruction: {0}\n  Waypoint: {1}\n  New location: {2}\n",
                    instruction,
                    waypointLocation,
                    ferryLocation
                )
            )
        ;
    }
}
