using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.IO;
using System.Linq;


namespace AoC
{
    public class Tiler
    {
        public void Lay(IEnumerable<TileToken> tiles)
        {
            var blackTiles = new Dictionary<Location, TileToken>();
            var tilesFlipped = 0;

            Console.WriteLine("\nLaying tiles:");
            foreach(var tile in tiles)
                FlipFile(tile);


            Console.WriteLine($"\nStats:\n  Tiles flipped: {tilesFlipped}\n  Black tiles: {blackTiles.Count}");
            Flip(blackTiles);
            return;


            void FlipFile(TileToken tile)
            {
                if(blackTiles.ContainsKey(tile.Location))
                    blackTiles.Remove(tile.Location);
                else
                    blackTiles.Add(tile.Location, tile.Flip());

                Console.WriteLine($"  {tile.Location}");
                tilesFlipped++;
            }
        }


        private void Flip(Dictionary<Location, TileToken> blackTiles)
        {
            Console.WriteLine("Simulating Days:");
            for(var day = 1; day <= 10; day++)
            {
                var whiteToFlip = WhiteTilesWithTwoBlackNeighbours(blackTiles).ToList();
                var blackToFlip = BlackTilesWithZeroThreeOrMoreBlackNeigbours(blackTiles).ToList();

                blackToFlip.ForEach(b => blackTiles.Remove(b));
                whiteToFlip.ForEach(w => blackTiles.Add(w, new TileToken(Colour.Black, w, new List<Direction>())));

                Console.WriteLine($"  Day {day}: {blackTiles.Count()}");
            }
        }

        private IEnumerable<Location> WhiteTilesWithTwoBlackNeighbours(Dictionary<Location, TileToken> blackTiles) =>
            from whiteTile in GetWhiteTiles(blackTiles.Keys)
            where CountBlackNeighbours(blackTiles, whiteTile) is 2
            select whiteTile
        ;

        private IEnumerable<Location> BlackTilesWithZeroThreeOrMoreBlackNeigbours(Dictionary<Location, TileToken> blackTiles) =>
            from blackTile in blackTiles
            where CountBlackNeighbours(blackTiles, blackTile.Key) is 0 or > 2
            select blackTile.Key
        ;

        private int CountBlackNeighbours(Dictionary<Location, TileToken> blackTiles, Location location) =>
            GetNeighbours(location).Where(n => blackTiles.ContainsKey(n)).Count()
        ;

        private IEnumerable<Location> GetNeighbours(Location location)
        {
            yield return new Location(location.Northing + 1, location.Easting + .5);    // ne
            yield return new Location(location.Northing,     location.Easting +  1);    // e
            yield return new Location(location.Northing - 1, location.Easting + .5);    // se
            yield return new Location(location.Northing - 1, location.Easting - .5);    // sw
            yield return new Location(location.Northing,     location.Easting -  1);    // w
            yield return new Location(location.Northing + 1, location.Easting - .5);    // nw
        }

        private IEnumerable<Location> GetWhiteTiles(IEnumerable<Location> blackTiles)
        {
            var maxWest = Math.Floor(blackTiles.Min(t => t.Easting) - 1.0);
            var maxEast = Math.Ceiling(blackTiles.Max(t => t.Easting) + 1.0);
            var maxNorth = blackTiles.Min(t => t.Northing) - 1.0;
            var maxSouth = blackTiles.Max(t => t.Northing) - 1.0;


            for(var northing = maxSouth; northing <= maxNorth; northing++)
                for(var easting = maxWest; easting <= maxEast; easting++)
                {
                    var offset = northing % 2 == 1 ? .5 : 0;
                    var location = new Location(northing, easting + offset);

                    if( ! blackTiles.Contains(location) )
                        yield return location;
                }
        }
    }
}
