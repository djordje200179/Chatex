using System.Net.Sockets;

namespace Chatex;
public class Client {
	private readonly TcpClient client;

	public Client(string address, ushort port) {
		client = new TcpClient(address, port);
	}
}