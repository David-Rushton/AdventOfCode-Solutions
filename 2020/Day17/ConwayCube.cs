namespace AoC
{
    public record ConwayCubePosition
    (
        int X,
        int Y,
        int Z
    )
    {
        public string Id => string.Format("{0}:{1}:{2}", this.X, this.Y, this.Z);
    };


    public record ConwayCube
    (
        ConwayCubePosition Position,
        bool IsActive
    )
    {
        public ConwayCube(int x, int y, int z, bool isActive)
            : this
            (
                new ConwayCubePosition(x, y, z),
                isActive
            )
        { }


        // Cube and it's position share a common id.
        // The id is its unique location (x, y and z address).
        public string Id => Position.Id;
    }
}
