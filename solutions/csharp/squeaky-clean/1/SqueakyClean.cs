using System.Text;

public static class Identifier
{
    public static string Clean(string identifier)
    {
        if (identifier == null) return null;
        var sb = new StringBuilder(identifier);
        for (int i = sb.Length - 1; i >= 0; i--) {
            if (sb[i] == ' ') {
                sb[i] = '_';
            } else if (sb[i] < ' ') {
                sb.Remove(i, 1);
                sb.Insert(i, "CTRL");
            } else if (sb[i] == '-') {
                sb.Remove(i, 1); // убираем дефис

                if (i < sb.Length) {
                    sb[i] = char.ToUpper(sb[i]); // первую букву следующего слова делаем заглавной
                }
            } else if (!char.IsLetter(sb[i])) {
                sb.Remove(i, 1);
            } else if (sb[i] >= 'α' && sb[i] <= 'ω') {
                sb.Remove(i, 1);
            }
            
        }

        return sb.ToString();
    }
}
