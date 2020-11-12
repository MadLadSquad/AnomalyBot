package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	// Token is your discord bot's token as a string
	Token string
)

func init() {

	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Starting the avanta!")
	fmt.Println("Configuring the avanta process!")
	fmt.Println("Recruiting the avantajii!")
	fmt.Println("We are ready to avantim people!")
	fmt.Println("Press CTRL+C to exit!")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	prefix := "&"
	content := strings.ToLower(m.Content)

	s.UpdateStreamingStatus(1, "&help | Click Watch For Epic Hacker Music!", "https://www.youtube.com/watch?v=t6KFfYdNPh8")

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.HasPrefix(content, prefix+"avantainfo") {
		embed := NewEmbed().
			SetTitle("Avanta Information").
			AddField("Avanta Description", `Avanta Software Foundation is a software company based in Gabrovo, Bulgaria. We believe in the "Аванта" which means something you get for free from somebody but you don't return it.`).
			AddField("Github:", "https://github.com/AvantaSoftwareFoundation").
			SetImage("https://raw.githubusercontent.com/AvantaSoftwareFoundation/AvantaSoftwareFoundationReadme/master/avanta.png").
			SetFooter(`We distribute free software, free as in "avanta"!`).
			SetColor(0xf1c40f).MessageEmbed

		s.ChannelMessageSendEmbed(m.ChannelID, embed)
	}

	if strings.HasPrefix(content, prefix+"help") {
		embed := NewEmbed().
			SetTitle("Help").
			AddField("&verify", "Verifies you by giving you the member and tester role.").
			AddField("&avantainfo", "Returns information about the AvantaSoftwareFoundation").
			AddField("&accinfo", "Returns info about your Anomaly Account").
			AddField("ping", "Returns pong.").
			AddField("pong", "Returns ping.").
			SetFooter(`Message delivered by Go™`).
			SetColor(0xf1c40f).MessageEmbed

		s.ChannelMessageSendEmbed(m.ChannelID, embed)
	}

	if strings.HasPrefix(content, prefix+"accinfo") {

		if m.Author.ID == "695224936985264219" {
			embed := NewEmbed().
				SetTitle("Anomaly Account Info!").
				AddField("Rank", "Private Rank 1").
				AddField("Player Type", "Creator of the game").
				AddField("Favourite Game Mode", "Data not available!").
				SetFooter(`Message delivered by Go™`).
				SetColor(0xf1c40f).MessageEmbed

			s.ChannelMessageSendEmbed(m.ChannelID, embed)
		} else if m.Author.ID == "648527310814707722" || m.Author.ID == "287966817383874560" || m.Author.ID == "324617194795171840" || m.Author.ID == "746389802521460736" || m.Author.ID == "165554929107206144" || m.Author.ID == "643205010292736020" || m.Author.ID == "330227754287104001" {
			embed := NewEmbed().
				SetTitle("Anomaly Account Info!").
				AddField("Rank", "Private Rank 1").
				AddField("Player Type", "Moderator").
				AddField("Favourite Game Mode", "Data not available!").
				SetFooter(`Message delivered by Go™`).
				SetColor(0xf1c40f).MessageEmbed

			s.ChannelMessageSendEmbed(m.ChannelID, embed)
		} else if m.Author.ID == "443392402007523328" || m.Author.ID == "407336011186831362" || m.Author.ID == "497125688848154626" || m.Author.ID == "278126141661773825" || m.Author.ID == "468843330076147733" || m.Author.ID == "349255625160785930" || m.Author.ID == "249886748908191744" {
			embed := NewEmbed().
				SetTitle("Anomaly Account Info!").
				AddField("Rank", "Private Rank 1").
				AddField("Player Type", "Pre-Alpha Tester").
				AddField("Favourite Game Mode", "Data not available!").
				SetFooter(`Message delivered by Go™`).
				SetColor(0xf1c40f).MessageEmbed

			s.ChannelMessageSendEmbed(m.ChannelID, embed)
		} else if m.Author.ID == "463959883046453249" || m.Author.ID == "622462671131901973" || m.Author.ID == "469999280422256669" || m.Author.ID == "295979779570532352" || m.Author.ID == "412577083915501578" || m.Author.ID == "708752218400227458" || m.Author.ID == "565589902465433708" || m.Author.ID == "687309713355833377" || m.Author.ID == "363329082316029952" || m.Author.ID == "213417506692399104" || m.Author.ID == "249523486828003328" || m.Author.ID == "604631268289937419" || m.Author.ID == "676763652367450112" || m.Author.ID == "696951411988103169" || m.Author.ID == "191254252402900993" || m.Author.ID == "520696535311188000" || m.Author.ID == "455806324131889162" || m.Author.ID == "615496288800342056" || m.Author.ID == "714060687072034876" || m.Author.ID == "411987480607195136" || m.Author.ID == "365859941292048384" || m.Author.ID == "266874291155501056" || m.Author.ID == "685848894923341858" {
			embed := NewEmbed().
				SetTitle("Anomaly Account Info!").
				AddField("Rank", "Private Rank 1").
				AddField("Player Type", "Alpha Tester").
				AddField("Favourite Game Mode", "Data not available!").
				SetFooter(`Message delivered by Go™`).
				SetColor(0xf1c40f).MessageEmbed

			s.ChannelMessageSendEmbed(m.ChannelID, embed)
		} else {
			embed := NewEmbed().
				SetTitle("Anomaly Account Info!").
				AddField("Rank", "Private Rank 1").
				AddField("Player Type", "Beta Tester").
				AddField("Favourite Game Mode", "Data not available!").
				SetFooter(`Message delivered by Go™`).
				SetColor(0xf1c40f).MessageEmbed
			s.ChannelMessageSendEmbed(m.ChannelID, embed)
		}
	}

	if strings.Contains(content, "&police") {
		s.ChannelMessageSend(m.ChannelID, "https://cdn.discordapp.com/attachments/676494184626126858/739786778327777380/unknown.png")
	}

	if strings.Contains(content, "&brazil") {
		s.ChannelMessageSend(m.ChannelID, "https://cdn.discordapp.com/attachments/676494184626126858/741964473522454538/unknown.png")
	}

	if strings.Contains(content, "&kfc") {
		s.ChannelMessageSend(m.ChannelID, "https://cdn.discordapp.com/attachments/209843216227368962/737819474043863172/video0.mp4")
	}

	if strings.Contains(content, "&epic-vid") {
		s.ChannelMessageSend(m.ChannelID, "Write &kfc!")
	}

	if strings.HasPrefix(content, prefix+"&cool-server") {
		s.ChannelMessageSend(m.ChannelID, "https://discord.gg/Fyejzu2")
	}

	if strings.HasPrefix(content, prefix+"verify") {
		s.GuildMemberRoleAdd(m.GuildID, m.Author.ID, "573049402428620810")
		s.GuildMemberRoleAdd(m.GuildID, m.Author.ID, "571288573781737483")
		s.GuildMemberRoleAdd(m.GuildID, m.Author.ID, "717037253691572327")
		s.GuildMemberRoleAdd(m.GuildID, m.Author.ID, "717037253292982321")
		s.ChannelMessageDelete(m.ChannelID, m.ID)

	}
	if strings.Contains(content, "avanta") || strings.Contains(content, "аванта") {
		s.ChannelMessageSend(m.ChannelID, "Аванта ли чух!?")
	}

	if strings.HasPrefix(content, prefix+"project-cursed") {
		s.ChannelMessageSend(m.ChannelID, "You mean Project Anomaly")
		s.ChannelMessageSend(m.ChannelID, "https://cdn.discordapp.com/attachments/747865814161686642/750771964666314753/unknown.png")
	}

	if strings.Contains(content, prefix+"mcdonalds") {
		s.ChannelMessageSend(m.ChannelID, "https://cdn.discordapp.com/attachments/758053551082831973/762254042293862440/Ronald_McDonald_insanity.mp4")

	}
}
