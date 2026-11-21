class RemoteControlCar
{
    // TODO: define the constructor for the 'RemoteControlCar' class
    private int speed;
    private int batteryDrain;
    private int distance;
    private int battery = 100;

    public int Speed() {
        return speed;
    }

    public int BatteryDrain() {
        return batteryDrain;
    }
    
    public RemoteControlCar(int speed, int batteryDrain) {
        this.speed = speed;
        this.batteryDrain = batteryDrain;
        this.distance = 0;
    }
    public bool BatteryDrained()
    {
        return battery < batteryDrain;
    }

    public int DistanceDriven()
    {
        return distance;
    }

    public void Drive()
    {
        if (battery < batteryDrain) return;

        distance += speed;
        battery -= batteryDrain;
    }
    
    public static RemoteControlCar Nitro()
    {
        return new RemoteControlCar(50, 4);
    }
}

class RaceTrack
{
    // TODO: define the constructor for the 'RaceTrack' class
    private int distance;
    public RaceTrack(int distance) {
        this.distance = distance;
    }
    public bool TryFinishTrack(RemoteControlCar car)
    {
        return (long)car.Speed() * 100 >= (long)distance * car.BatteryDrain();
    }
}
