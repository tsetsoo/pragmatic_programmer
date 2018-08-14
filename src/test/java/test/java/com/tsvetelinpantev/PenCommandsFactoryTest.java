package test.java.com.tsvetelinpantev;


import com.tsvetelinpantev.PenCommandsFactory;
import org.junit.Assert;
import org.junit.Test;
import org.junit.jupiter.api.Assertions;

public class PenCommandsFactoryTest {

    private PenCommandsFactory penCommandsFactory = PenCommandsFactory.init();

    @Test
    public void testSelectPenCommand() {
        String commandResult = penCommandsFactory.execute("P", "2");
        Assert.assertEquals("Unexpected select pen command result", "select pen 2", commandResult);
    }

    @Test
    public void testPenDownCommand() {
        String commandResult = penCommandsFactory.execute("D");
        Assert.assertEquals("Unexpected pen down command result", "pen down", commandResult);
    }

    @Test
    public void testDrawWestCommand() {
        String commandResult = penCommandsFactory.execute("W", "1");
        Assert.assertEquals("Unexpected draw west command result", "draw west 1cm", commandResult);
    }

    @Test
    public void testDrawEastCommand() {
        String commandResult = penCommandsFactory.execute("E", "2");
        Assert.assertEquals("Unexpected draw east command result", "draw east 2cm", commandResult);
    }

    @Test
    public void testDrawNorthCommand() {
        String commandResult = penCommandsFactory.execute("N", "3");
        Assert.assertEquals("Unexpected draw north command result", "draw north 3cm", commandResult);
    }

    @Test
    public void testDrawSouthCommand() {
        String commandResult = penCommandsFactory.execute("S", "10");
        Assert.assertEquals("Unexpected draw south command result", "draw south 10cm", commandResult);
    }

    @Test
    public void testPenUpCommand() {
        String commandResult = penCommandsFactory.execute("U");
        Assert.assertEquals("Unexpected pen up command result", "pen up", commandResult);
    }

    @Test
    public void nonExistentCommand() {
        Assertions.assertThrows(IllegalArgumentException.class, () -> penCommandsFactory.execute("A"));
    }

    @Test
    public void testSelectPenCommandInvalidArguments() {
        Assertions.assertThrows(IllegalArgumentException.class, () -> penCommandsFactory.execute("P", "2", "2"));
    }

    @Test
    public void testPenDownCommandInvalidArguments() {
        Assertions.assertThrows(IllegalArgumentException.class, () -> penCommandsFactory.execute("D", "2"));
    }

    @Test
    public void testDrawWestCommandInvalidArguments() {
        Assertions.assertThrows(IllegalArgumentException.class, () -> penCommandsFactory.execute("W", "2", "2"));
    }

    @Test
    public void testDrawEastCommandInvalidArguments() {
        Assertions.assertThrows(IllegalArgumentException.class, () -> penCommandsFactory.execute("E", "2", "2"));
    }

    @Test
    public void testDrawNorthCommandInvalidArguments() {
        Assertions.assertThrows(IllegalArgumentException.class, () -> penCommandsFactory.execute("N", "2", "2"));
    }

    @Test
    public void testDrawSouthCommandInvalidArguments() {
        Assertions.assertThrows(IllegalArgumentException.class, () -> penCommandsFactory.execute("S", "2", "2"));
    }

    @Test
    public void testPenUpCommandInvalidArguments() {
        Assertions.assertThrows(IllegalArgumentException.class, () -> penCommandsFactory.execute("U", "2"));
    }
}
