using System;
using System.Globalization;

static class Appointment
{
    public static DateTime Schedule(string appointmentDateDescription)
    {
        return DateTime.Parse(appointmentDateDescription, CultureInfo.InvariantCulture);
    }

    public static bool HasPassed(DateTime appointmentDate)
    {
        return appointmentDate < DateTime.Now;
    }

    public static bool IsAfternoonAppointment(DateTime appointmentDate)
    {
        TimeSpan time = appointmentDate.TimeOfDay;
        return time >= new TimeSpan(12, 0, 0) && time < new TimeSpan(18, 0, 0);
        /*
        int hour = appointmentDate.Hour;
        return hour >= 12 && hour < 18;
        */
    }

    public static string Description(DateTime appointmentDate)
    {
        return $"You have an appointment on {appointmentDate.ToString("M/d/yyyy h:mm:ss tt", CultureInfo.InvariantCulture)}.";
    }

    public static DateTime AnniversaryDate()
    {
        return new DateTime(DateTime.Now.Year, 9, 15);
    }
}
