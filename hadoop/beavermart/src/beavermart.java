import java.util.Scanner;
import java.util.regex.Matcher;
import java.util.regex.Pattern;
import java.io.BufferedWriter;
import java.io.File;
import java.io.IOException;
import java.io.FileOutputStream;
import java.io.FileNotFoundException;
import java.io.OutputStreamWriter;

import org.apache.hadoop.conf.Configuration;
import org.apache.hadoop.fs.Path;
import org.apache.hadoop.mapreduce.Mapper;
import org.apache.hadoop.mapreduce.Reducer;
import org.apache.hadoop.io.IntWritable;
import org.apache.hadoop.io.LongWritable;
import org.apache.hadoop.io.Text;
import org.apache.hadoop.util.GenericOptionsParser;
import org.apache.hadoop.mapreduce.Job;
import org.apache.hadoop.mapreduce.lib.input.FileInputFormat;
import org.apache.hadoop.mapreduce.lib.output.FileOutputFormat;

public class Beavermart {
	static int unique=0;
	static int total=0;

	public static void main(String [] args) throws Exception
	{
		Configuration c=new Configuration();
		String[] files=new GenericOptionsParser(c,args).getRemainingArgs();
		Path input=new Path(files[0]);
		Path output=new Path(files[1]);
		Job j = Job.getInstance(c,"wordcount");
		j.setJarByClass(Beavermart.class);
		j.setMapperClass(MapForMp2.class);
		j.setReducerClass(ReduceForMp2.class);
		j.setOutputKeyClass(Text.class);
		j.setOutputValueClass(IntWritable.class);
		FileInputFormat.addInputPath(j, input);
		FileOutputFormat.setOutputPath(j, output);
		j.waitForCompletion(true);
		System.out.println("unique word pairs: "+unique);
		System.out.println("total word pairs: "+total);	
		System.exit(0);
	}
	2 jum

	public static class MapForMp2 extends Mapper<LongWritable, Text, Text, IntWritable>{
		public void map(LongWritable key, Text value, Context con) throws IOException, InterruptedException
		{
			String line = value.toString();
			String[] words=line.split(",");
			int len = words.length;
			for(int i = 0; i < len; i++)
			{
				for(int j = i+1; j < len; j++)
				{
					String str1 = words[i].trim();
					String str2 = words[j].trim();
					Text outputKey = new Text();

					if (str1.compareTo(str2) < 0)
						outputKey.set("(" + str1 + ", " + str2 + ")");
					else
						outputKey.set("(" + str2 + ", " + str1 + ")");
					IntWritable outputValue = new IntWritable(1);
					con.write(outputKey, outputValue);
				}
			}
		}
	}


	public static class ReduceForMp2 extends Reducer<Text, IntWritable, Text, IntWritable>
	{
		public void reduce(Text word, Iterable<IntWritable> values, Context con) throws IOException, InterruptedException
		{
			int sum = 0;
			int k=0;
			for(IntWritable value : values)
			{
				sum += value.get();
				if(k==0) {
					unique+=1;
				}
				k=1;
				total+=1;
			}
			con.write(word, new IntWritable(sum));
		}
	}
}