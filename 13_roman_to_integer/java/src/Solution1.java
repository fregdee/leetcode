import java.util.HashMap;
import java.util.Map;

public class Solution1 {

    // TODO: なぜ values 変数の型は Map なのか調べる
    static Map<String, Integer> values = new HashMap<>();

    static {
        values.put("M", 1000);
        values.put("D", 500);
        values.put("C", 100);
        values.put("L", 50);
        values.put("X", 10);
        values.put("V", 5);
        values.put("I", 1);
    }

    public static void main(String[] args) {
        String s = "CMXCIV";
        int result = romanToInt(s);
        System.out.println(result);
    }

    static int romanToInt(String s) {
        int sum = 0;
        int i = 0;
        while (i < s.length()) {
            // TODO: 特定の1文字だけ取得する方法まとめる ( s.substring(i) じゃダメ？ )
            String currentString = s.substring(i, i+1);
            int currentValue = values.get(currentString);
            int nextValue = 0;

            if (i+1 < s.length()) {
                String nextString = s.substring(i+1, i+2);
                nextValue = values.get(nextString);
            }

            if (currentValue < nextValue) {
                sum += (nextValue - currentValue);
                i += 2;
            } else {
                sum += currentValue;
                i += 1;
            }
        }
        return sum;
    }
}
