//java libs
import java.io.*;
import java.io.IOException;
import java.util.Scanner;
import java.util.regex.Matcher;
import java.util.regex.Pattern;
import java.net.URI;

//hadoop libs
import org.apache.hadoop.conf.Configuration;
import org.apache.hadoop.fs.*;
import org.apache.hadoop.io.IntWritable;
import org.apache.hadoop.io.LongWritable;
import org.apache.hadoop.io.*;
import org.apache.hadoop.conf.*;
import org.apache.hadoop.util.GenericOptionsParser;
import org.apache.hadoop.mapreduce.Job;
import org.apache.hadoop.mapreduce.Mapper;
import org.apache.hadoop.mapreduce.Reducer;
import org.apache.hadoop.mapreduce.lib.input.FileInputFormat;
import org.apache.hadoop.mapreduce.lib.output.FileOutputFormat;


/*mainclass */ 
public class wordcount {
	static int totalcount=0; 
	static int uni=0;
	static int srno= 1;
	public static void main(String [] args) throws Exception
	{
		Configuration c=new Configuration();
		String[] files=new GenericOptionsParser(c,args).getRemainingArgs();
		Path input=new Path(files[0]);
		Path output=new Path(files[1]);
		Job j=new Job(c,"wordcount");
		j.setJarByClass(wordcount.class);
		j.setMapperClass(MapForwordcount.class);
		j.setReducerClass(ReduceForwordcount.class);
		j.setOutputKeyClass(Text.class);
		j.setOutputValueClass(IntWritable.class);
		FileInputFormat.addInputPath(j, input);
		FileOutputFormat.setOutputPath(j, output);
		j.waitForCompletion(true);
		
		System.out.println("Unique Words: "+uni);
		System.out.println("Total counted Words: "+totalcount);

	}
//Mapper -- defines input and output key and input and output datatypes
	public static class MapForwordcount extends Mapper<LongWritable, Text, Text, IntWritable>{
		/* map function --takes the text and creates key value pairs for each word */
		public void map(LongWritable key, Text value, Context con) throws InterruptedException, IOException 
		{
			String line = value.toString(); //convert whole text to string 
			String[] words=line.split(" "); // split strings into individual words
			for(String word: words )
			{
				String result = word.replaceAll("[-+.^:,]","");
				Text outputKey = new Text(result.toLowerCase().trim());
				IntWritable outputValue = new IntWritable(1);
				con.write(outputKey, outputValue);
			}
		}
	}
//Reducer -- 
	public static class ReduceForwordcount extends Reducer<Text, IntWritable, Text, IntWritable>
	{
		public void reduce(Text word, Iterable<IntWritable> values, Context con) throws InterruptedException, IOException
		{
			int add = 0;
			int x=0;
			String st=""+srno+") "+word.toString();
			Text word_new=new Text();
			word_new.set(st);
			for(IntWritable value : values)
			{
				add += value.get();
				if(x==0) {
					uni+=1;
				}
				totalcount+=1;
				x=1;
			}
			con.write(word_new, new IntWritable(add));
			srno++;
		}
		

	}

}
