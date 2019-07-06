# News-GO
A RSS Feed server for the **[PSP News Channel](https://github.com/PSPConnect64/News-Channel-PSP)**, which will make it very easy to create and serve PSP-friendly RSS files.

So, how does it work?
Basically, this could have been done on the PSP itself, but hey, we're retards.
Connect via a TCP socket and send the topic, you'll receive the XML directly. The connection won't be closed by the server, you can send as many topic request as you want. But please - don't send a DDoS attack, we're not that stupid, are we?