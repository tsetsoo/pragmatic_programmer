class Blender {

    private boolean isFull;
    private int speed;

    public Blender() {
        this.isFull = false;
        this.speed = 0;
    }

    public void setSpeed(final int x) {
        if (Math.abs(getSpeed() - x) > 1) {
            throw new BlenderException("Cannot change speed by more than one");
        }
        this.speed = x;
    }

    public int getSpeed() {
        return this.speed;
    }

    public void fill() {
        if (this.isFull) {
            throw new BlenderException("Cannot fill full blender");
        }
        this.isFull = true;
    }

    public void empty() {
        if (!this.isFull) {
            throw new BlenderException("Cannot empty not full blender");
        }
        this.isFull = false;
    }
}

class BlenderException extends RuntimeException {
    BlenderException(String message) {
        super(message);
    }
}

public class BlenderTest {
    public static void main(String[] args) {
        testSetSpeed1When0();
        testSetSpeed2When0();
        testSetSpeed0When1();
        testSetSpeed0When2();
        testSetSpeed11When10();
        testFillWhenNotFull();
        testEmptyWhenFull();
        testFillWhenFull();
        testEmptyWhenNotFull();
        System.out.println("success");
    }

    private static void testSetSpeed1When0() {
        Blender blender = new Blender();
        blender.setSpeed(1);
        assert blender.getSpeed() == 1;
    }

    private static void testSetSpeed2When0() {
        Blender blender = new Blender();
        try {
            blender.setSpeed(2);
            assert false;
        } catch (BlenderException e) {
        }
    }

    private static void testSetSpeed0When1() {
        Blender blender = new Blender();
        blender.setSpeed(1);
        blender.setSpeed(0);
        assert blender.getSpeed() == 0;
    }

    private static void testSetSpeed0When2() {
        Blender blender = new Blender();
        blender.setSpeed(1);
        blender.setSpeed(2);
        try {
            blender.setSpeed(0);
            assert false;
        } catch (BlenderException e) {
        }
    }

    private static void testSetSpeed11When10() {
        Blender blender = new Blender();
        for (int i = 1; i < 11; i++) {
            blender.setSpeed(i);
        }
        try {
            blender.setSpeed(11);
            assert false;
        } catch (BlenderException e) {
        }
    }

    private static void testFillWhenNotFull() {
        Blender blender = new Blender();
        blender.fill();
    }

    private static void testEmptyWhenFull() {
        Blender blender = new Blender();
        blender.fill();
        blender.empty();
    }

    private static void testFillWhenFull() {
        Blender blender = new Blender();
        blender.fill();
        try {
            blender.fill();
            assert false;
        } catch (BlenderException e) {
        }
    }

    private static void testEmptyWhenNotFull() {
        Blender blender = new Blender();
        try {
            blender.empty();
            assert false;
        } catch (BlenderException e) {
        }
    }
}
