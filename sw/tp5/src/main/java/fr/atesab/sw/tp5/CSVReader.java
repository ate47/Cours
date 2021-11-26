package fr.atesab.sw.tp5;

import java.io.BufferedReader;
import java.io.File;
import java.io.FileReader;
import java.nio.charset.Charset;
import java.util.function.Consumer;
import java.util.function.Function;
import java.util.stream.Stream;

import org.apache.commons.codec.Charsets;

public class CSVReader<T> {
    private File file;
    private Function<String[], T> reading;

    public CSVReader(String file, Function<String[], T> reading) {
        this(new File(file), reading);
    }

    public CSVReader(File file, Function<String[], T> reading) {
        this.file = file;
        this.reading = reading;
    }

    public void readFile(Consumer<T> action) {
        readFile(-1, action);
    }

    public void readFile(int maxLine, Consumer<T> action) {
        try (var r = new BufferedReader(new FileReader(file, Charsets.UTF_8))) {
            r.readLine(); // header
            String line;
            var i = 0;
            while ((line = r.readLine()) != null) {
                if (line.isEmpty())
                    continue;

                String[] raw;
                if (line.charAt(line.length() - 1) == ',') {
                    raw = (line + ' ').split(",");
                    raw[raw.length - 1] = "";
                } else {
                    raw = line.split(",");
                }

                if (maxLine != -1 && i++ >= maxLine) {
                    break;
                }

                var t = reading.apply(raw);
                action.accept(t);
            }
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
}
