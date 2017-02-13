package cp;

public class StringTest{
	public static void main(String[] args) {
		String s1 = "pq";
		String s2 = "pq";

		System.out.println(s1 == s2);//true
		String x = "q";
		String s3 = "p" + x;
		System.out.println(s3 == s1);//false

		s3 = s3.intern();
		System.out.println(s3 == s1);//true
	}
}