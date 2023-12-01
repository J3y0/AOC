import java.io.*;
import java.util.HashMap;

public class App {
    public static void main(String[] args) throws Exception {
        // Read file
        String filename = "./day6.txt";
        FileReader in = new FileReader(filename);
		BufferedReader bin = new BufferedReader(in);

		String data = "";
		while(bin.ready()) {
			data = bin.readLine();
		}
		bin.close();

        // Part 1: detecting 4 distinct characters
        int i = 2;
        String extract = "jj";
        while ((same(extract)) && (i < data.length())) {
            i ++;
            extract = data.substring(i - 3, i + 1);
        }
        System.out.println(extract + " " + (i + 1));

        // Part 2: same but for 14
        int i2 = 12;
        String extract2 = "jj";
        while ((same(extract2)) && (i2 < data.length())) {
            i2 ++;
            extract2 = data.substring(i2 - 13, i2 + 1);
        }
        System.out.println(extract2 + " " + (i2 + 1));
    }

    static public boolean same(String extract) {
        // True if a letter is more than one time in the extract
        HashMap<Character, Integer> letters = new HashMap<>();

        for (int i = 0; i < extract.length(); i++) {
            char c = extract.charAt(i);
            if (letters.containsKey(c)) {
                letters.put(c, letters.get(c) + 1);
            } else {
                letters.put(c, 1);
            }
        }
        boolean result = false;
        for (Character c: letters.keySet()) {
            if (letters.get(c) > 1) {
                result = true;
                break;
            }
        }

        return result;
    }
}
