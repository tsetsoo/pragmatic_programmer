public class SimpleAssertion {

    public void assertion(boolean condition) {
        if (!condition) {
            Thread.dumpStack();
            throw new RuntimeException("Assertion failed");
        }
    }

    public static void main(String[] args) {
        SimpleAssertion asser = new SimpleAssertion();
        asser.assertion(1 == 0);
    }
}