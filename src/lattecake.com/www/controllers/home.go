package controllers

import (
	"lattecake.com/www/models"
	"log"
)

type HomeController struct {
	BaseController
}

func (c *HomeController) Get() {

	ch := make(chan models.Posts)
	starCh := make(chan models.Recommend)

	travelCh := make(chan models.Posts)
	starTravelCh := make(chan []models.Recommend)
	starMusicCh := make(chan []models.Recommend)
	musicCh := make(chan models.Posts)

	var learns models.Posts
	var starPost models.Recommend

	var travelPosts models.Posts
	var starTravelPosts []models.Recommend
	var starMusic []models.Recommend
	var musics []models.Post

	go func() {
		star, err := (&models.Recommend{}).StarPost()
		if err != nil {
			log.Println(err)
		}
		starCh <- star
	}()
	go func() {
		star, err := (&models.Recommend{}).HomePost(3, 3)
		if err != nil {
			log.Println(err)
		}
		starTravelCh <- star
	}()
	go func() {
		star, err := (&models.Recommend{}).HomePost(4, 3)
		if err != nil {
			log.Println(err)
		}
		starMusicCh <- star
	}()

	go func() {
		posts, err := (&models.Post{}).LastPosts(2, 3)
		if err != nil {
			log.Println(err)
		}
		travelCh <- posts
	}()

	go func() {
		posts, err := (&models.Post{}).LastPosts(4, 3)
		if err != nil {
			log.Println(err)
		}
		musicCh <- posts
	}()

	go func() {
		learn, _ := (&models.Post{}).LastPosts(1, 3)
		ch <- learn
	}()




	learns = <-ch
	starPost = <-starCh

	travelPosts = <-travelCh
	starTravelPosts = <-starTravelCh
	starMusic = <-starMusicCh
	musics = <-musicCh

	close(ch)
	close(starCh)

	close(travelCh)
	close(starTravelCh)


	c.Data["Learns"] = learns
	c.Data["StarPost"] = starPost
	c.Data["TravelPosts"] = travelPosts
	c.Data["TravelRecommend"] = starTravelPosts
	c.Data["MusicRecommend"] = starMusic
	c.Data["Musics"] = musics

	c.TplName = "home/index.tpl"
}
