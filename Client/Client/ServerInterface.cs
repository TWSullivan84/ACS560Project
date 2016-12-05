using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Net;
using System.Net.Sockets;

namespace ServerSocket
{
    public static class ServerInterface
    {
        public static void sendToServer(string gameData)
        {
            try
            {
                IPHostEntry ipHostInfo = Dns.GetHostEntry("ec2-35-163-106-205.us-west-2.compute.amazonaws.com");
                IPAddress ipAddress = ipHostInfo.AddressList[0]; 
                IPEndPoint remoteEP = new IPEndPoint(ipAddress, 11337);

                Socket sender = new Socket(AddressFamily.InterNetwork, SocketType.Stream, ProtocolType.Tcp);

                sender.Connect(remoteEP);

                Console.WriteLine("Socket connected to {0}", sender.RemoteEndPoint.ToString());

                byte[] msg = Encoding.ASCII.GetBytes(gameData);
                int bytesSent = sender.Send(msg);

                sender.Shutdown(SocketShutdown.Both);
                sender.Close();
            }
            catch(Exception e)
            {
                Console.WriteLine(e.Message);
            }
        }        
    }
}
