package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type gameBag struct {
	Game Game
}

type Game struct {
	ID                               string   `json:"id"`
	Gold                             string   `json:"gold"`
	Page                             string   `json:"page"`
	Notes                            string   `json:"notes"`
	Players                          []Player `json:"players"`
	FlagBouncerBounced               bool     `json:"flag-bouncer-bounced"`
	FlagCulpritIdentified            bool     `json:"flag-culprit-identified"`
	FlagSlipping                     bool     `json:"flag-slipping"`
	FlagContractTerminated           bool     `json:"flag-contract-terminated"`
	FlagPrizedPosession              bool     `json:"flag-prized-posession"`
	FlagLegWound                     bool     `json:"flag-leg-wound"`
	FlagSoulJudgement                bool     `json:"flag-soul-judgement"`
	FlagFamilialDispute              bool     `json:"flag-familial-dispute"`
	FlagExperiencedGuide             bool     `json:"flag-experienced-guide"`
	FlagALonelySurvivor              bool     `json:"flag-a-lonely-survivor"`
	FlagMarked                       bool     `json:"flag-marked"`
	FlagHappyFugitives               bool     `json:"flag-happy-fugitives"`
	FlagHeartbrokenExile             bool     `json:"flag-heartbroken-exile"`
	FlagJudgementOfJudas             bool     `json:"flag-judgement-of-judas"`
	FlagThePatelDaughter             bool     `json:"flag-the-patel-daughter"`
	FlagEliasGratitude               bool     `json:"flag-elias-gratitude"`
	FlagAnEndToMadness               bool     `json:"flag-an-end-to-madness"`
	FlagNewAmbitions                 bool     `json:"flag-new-ambitions"`
	FlagAvoidingCollapse             bool     `json:"flag-avoiding-collapse"`
	FlagAHusbandsDuty                bool     `json:"flag-a-husbands-duty"`
	FlagASonsLove                    bool     `json:"flag-a-sons-love"`
	FlagKeyDefeat                    bool     `json:"flag-key-defeat"`
	FlagRested                       bool     `json:"flag-rested"`
	FlagBookOfGehenna                bool     `json:"flag-book-of-gehenna"`
	FlagFallenFriend                 bool     `json:"flag-fallen-friend"`
	FlagLastOfHerCoterie             bool     `json:"flag-last-of-her-coterie"`
	FlagSeedsOfHerPeople             bool     `json:"flag-seeds-of-her-people"`
	FlagGrievousWounds               bool     `json:"flag-grievous-wounds"`
	FlagTheOnlyWay                   bool     `json:"flag-the-only-way"`
	FlagIrreconcilableDifferences    bool     `json:"flag-irreconcilable-differences"`
	FlagReluctantPartnership         bool     `json:"flag-reluctant-partnership"`
	FlagMassacred                    bool     `json:"flag-massacred"`
	FlagCoffinBuddy                  bool     `json:"flag-coffin-buddy"`
	FlagCapableEnough                bool     `json:"flag-capable-enough"`
	FlagGoodRiddance                 bool     `json:"flag-good-riddance"`
	FlagTimeIsOfTheEssence           bool     `json:"flag-time-is-of-the-essence"`
	FlagAllWeCanGet                  bool     `json:"flag-all-we-can-get"`
	FlagButchered                    bool     `json:"flag-butchered"`
	FlagSacrifice                    bool     `json:"flag-sacrifice"`
	FlagAHelpingHand                 bool     `json:"flag-a-helping-hand"`
	FlagANewOwner                    bool     `json:"flag-a-new-owner"`
	FlagAGreaterCause                bool     `json:"flag-a-greater-cause"`
	FlagImprisonedAnointed           bool     `json:"flag-imprisoned-anointed"`
	AchievementNotQuiteUnkillable    bool     `json:"achievement-not-quite-unkillable"`
	AchievementDeathFromAbove        bool     `json:"achievement-death-from-above"`
	AchievementProofOfProwess        bool     `json:"achievement-proof-of-prowess"`
	AchievementLikeABlur             bool     `json:"achievement-like-a-blur"`
	AchievementSupremeMobility       bool     `json:"achievement-supreme-mobility"`
	AchievementSickleSlaughter       bool     `json:"achievement-sickle-slaughter"`
	AchievementRainingLoot           bool     `json:"achievement-raining-loot"`
	AchievementHealthyWealthyAndWise bool     `json:"achievement-healthy-wealthy-and-wise"`
	AchievementSuccessBreedsSuccess  bool     `json:"achievement-success-breeds-success"`
	AchievementZapAGap               bool     `json:"achievement-zap-a-gap"`
	AchievementPoppetBeGone          bool     `json:"achievement-poppet-be-gone"`
	AchievementToughWurm             bool     `json:"achievement-tough-wurm"`
	AchievementTentacleCollector     bool     `json:"achievement-tentacle-collector"`
	AchievementTorturedMortal        bool     `json:"achievement-tortured-mortal"`
	UpgradeMasterwork                bool     `json:"upgrade-masterwork"`
	UpgradeImbued                    bool     `json:"upgrade-imbued"`
	UpgradeDevastating               bool     `json:"upgrade-devastating"`
	UpgradeOtherworldly              bool     `json:"upgrade-otherworldly"`
	UpgradeReinforced                bool     `json:"upgrade-reinforced"`
	UpgradeBulky                     bool     `json:"upgrade-bulky"`
	UpgradeElegant                   bool     `json:"upgrade-elegant"`
	UpgradeRefined                   bool     `json:"upgrade-refined"`
	UpgradeEtherium                  bool     `json:"upgrade-etherium"`
	UpgradeSentient                  bool     `json:"upgrade-sentient"`
	UpgradeScrying                   bool     `json:"upgrade-scrying"`
	UpgradeDark                      bool     `json:"upgrade-dark"`
	UpgradeGlowing                   bool     `json:"upgrade-glowing"`
	UpgradeShimmering                bool     `json:"upgrade-shimmering"`
	UpgradeEnchanted                 bool     `json:"upgrade-enchanted"`
}

type Player struct {
	Name         string `json:"name"`
	XP           string `json:"xp"`
	Damage       string `json:"damage"`
	Injured      bool   `json:"injured"`
	Unselectable bool   `json:"unselectable"`
}

func (s *Server) gameHandler(w http.ResponseWriter, r *http.Request) {
	bag := gameBag{}

	id := r.PathValue("id")
	if len(id) == 0 {
		errorResponse(w, http.StatusBadRequest, errGameIDRequired)
		return
	}

	cachePath := "game:" + id
	err := s.cache.Get(r.Context(), cachePath, &bag.Game)
	if err != nil {
		// errorResponse(w, http.StatusNotFound, err)
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	template := "game.gohtml"
	renderHtml(w, http.StatusOK, template, bag)
}

func (s *Server) saveGameHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if len(id) == 0 {
		errorResponse(w, http.StatusBadRequest, errGameIDRequired)
		return
	}

	cachePath := "game:" + id
	if found, _ := s.cache.Has(r.Context(), cachePath); !found {
		errorResponse(w, http.StatusNotFound, errGameNotFound)
		return
	}

	game := Game{
		ID:                               id,
		Gold:                             r.FormValue("gold"),
		Page:                             r.FormValue("page"),
		Notes:                            r.FormValue("notes"),
		FlagBouncerBounced:               r.FormValue("flag-bouncer-bounced") == "on",
		FlagCulpritIdentified:            r.FormValue("flag-culprit-identified") == "on",
		FlagSlipping:                     r.FormValue("flag-slipping") == "on",
		FlagContractTerminated:           r.FormValue("flag-contract-terminated") == "on",
		FlagPrizedPosession:              r.FormValue("flag-prized-posession") == "on",
		FlagLegWound:                     r.FormValue("flag-leg-wound") == "on",
		FlagSoulJudgement:                r.FormValue("flag-soul-judgement") == "on",
		FlagFamilialDispute:              r.FormValue("flag-familial-dispute") == "on",
		FlagExperiencedGuide:             r.FormValue("flag-experienced-guide") == "on",
		FlagALonelySurvivor:              r.FormValue("flag-a-lonely-survivor") == "on",
		FlagMarked:                       r.FormValue("flag-marked") == "on",
		FlagHappyFugitives:               r.FormValue("flag-happy-fugitives") == "on",
		FlagHeartbrokenExile:             r.FormValue("flag-heartbroken-exile") == "on",
		FlagJudgementOfJudas:             r.FormValue("flag-judgement-of-judas") == "on",
		FlagThePatelDaughter:             r.FormValue("flag-the-patel-daughter") == "on",
		FlagEliasGratitude:               r.FormValue("flag-elias-gratitude") == "on",
		FlagAnEndToMadness:               r.FormValue("flag-an-end-to-madness") == "on",
		FlagNewAmbitions:                 r.FormValue("flag-new-ambitions") == "on",
		FlagAvoidingCollapse:             r.FormValue("flag-avoiding-collapse") == "on",
		FlagAHusbandsDuty:                r.FormValue("flag-a-husbands-duty") == "on",
		FlagASonsLove:                    r.FormValue("flag-a-sons-love") == "on",
		FlagKeyDefeat:                    r.FormValue("flag-key-defeat") == "on",
		FlagRested:                       r.FormValue("flag-rested") == "on",
		FlagBookOfGehenna:                r.FormValue("flag-book-of-gehenna") == "on",
		FlagFallenFriend:                 r.FormValue("flag-fallen-friend") == "on",
		FlagLastOfHerCoterie:             r.FormValue("flag-last-of-her-coterie") == "on",
		FlagSeedsOfHerPeople:             r.FormValue("flag-seeds-of-her-people") == "on",
		FlagGrievousWounds:               r.FormValue("flag-grievous-wounds") == "on",
		FlagTheOnlyWay:                   r.FormValue("flag-the-only-way") == "on",
		FlagIrreconcilableDifferences:    r.FormValue("flag-irreconcilable-differences") == "on",
		FlagReluctantPartnership:         r.FormValue("flag-reluctant-partnership") == "on",
		FlagMassacred:                    r.FormValue("flag-massacred") == "on",
		FlagCoffinBuddy:                  r.FormValue("flag-coffin-buddy") == "on",
		FlagCapableEnough:                r.FormValue("flag-capable-enough") == "on",
		FlagGoodRiddance:                 r.FormValue("flag-good-riddance") == "on",
		FlagTimeIsOfTheEssence:           r.FormValue("flag-time-is-of-the-essence") == "on",
		FlagAllWeCanGet:                  r.FormValue("flag-all-we-can-get") == "on",
		FlagButchered:                    r.FormValue("flag-butchered") == "on",
		FlagSacrifice:                    r.FormValue("flag-sacrifice") == "on",
		FlagAHelpingHand:                 r.FormValue("flag-a-helping-hand") == "on",
		FlagANewOwner:                    r.FormValue("flag-a-new-owner") == "on",
		FlagAGreaterCause:                r.FormValue("flag-a-greater-cause") == "on",
		FlagImprisonedAnointed:           r.FormValue("flag-imprisoned-anointed") == "on",
		AchievementNotQuiteUnkillable:    r.FormValue("achievement-not-quite-unkillable") == "on",
		AchievementDeathFromAbove:        r.FormValue("achievement-death-from-above") == "on",
		AchievementProofOfProwess:        r.FormValue("achievement-proof-of-prowess") == "on",
		AchievementLikeABlur:             r.FormValue("achievement-like-a-blur") == "on",
		AchievementSupremeMobility:       r.FormValue("achievement-supreme-mobility") == "on",
		AchievementSickleSlaughter:       r.FormValue("achievement-sickle-slaughter") == "on",
		AchievementRainingLoot:           r.FormValue("achievement-raining-loot") == "on",
		AchievementHealthyWealthyAndWise: r.FormValue("achievement-healthy-wealthy-and-wise") == "on",
		AchievementSuccessBreedsSuccess:  r.FormValue("achievement-success-breeds-success") == "on",
		AchievementZapAGap:               r.FormValue("achievement-zap-a-gap") == "on",
		AchievementPoppetBeGone:          r.FormValue("achievement-poppet-be-gone") == "on",
		AchievementToughWurm:             r.FormValue("achievement-tough-wurm") == "on",
		AchievementTentacleCollector:     r.FormValue("achievement-tentacle-collector") == "on",
		AchievementTorturedMortal:        r.FormValue("achievement-tortured-mortal") == "on",
		UpgradeMasterwork:                r.FormValue("upgrade-masterwork") == "on",
		UpgradeImbued:                    r.FormValue("upgrade-imbued") == "on",
		UpgradeDevastating:               r.FormValue("upgrade-devastating") == "on",
		UpgradeOtherworldly:              r.FormValue("upgrade-otherworldly") == "on",
		UpgradeReinforced:                r.FormValue("upgrade-reinforced") == "on",
		UpgradeBulky:                     r.FormValue("upgrade-bulky") == "on",
		UpgradeElegant:                   r.FormValue("upgrade-elegant") == "on",
		UpgradeRefined:                   r.FormValue("upgrade-refined") == "on",
		UpgradeEtherium:                  r.FormValue("upgrade-etherium") == "on",
		UpgradeSentient:                  r.FormValue("upgrade-sentient") == "on",
		UpgradeScrying:                   r.FormValue("upgrade-scrying") == "on",
		UpgradeDark:                      r.FormValue("upgrade-dark") == "on",
		UpgradeGlowing:                   r.FormValue("upgrade-glowing") == "on",
		UpgradeShimmering:                r.FormValue("upgrade-shimmering") == "on",
		UpgradeEnchanted:                 r.FormValue("upgrade-enchanted") == "on",
	}

	for i := 0; i < 6; i++ {
		game.Players = append(game.Players, Player{
			Name:         r.FormValue(fmt.Sprintf("name-%d", i)),
			XP:           r.FormValue(fmt.Sprintf("xp-%d", i)),
			Damage:       r.FormValue(fmt.Sprintf("damage-%d", i)),
			Injured:      r.FormValue(fmt.Sprintf("injured-%d", i)) == "on",
			Unselectable: r.FormValue(fmt.Sprintf("unselectable-%d", i)) == "on",
		})
	}

	err := s.cache.Set(r.Context(), cachePath, &game, time.Hour*24*90)
	if err != nil {
		errorResponse(w, http.StatusInternalServerError, err)
		return
	}

	http.Redirect(w, r, "/game/"+id, http.StatusFound)
}

func (s *Server) newGameHandler(w http.ResponseWriter, r *http.Request) {
	id := s.buildGameKey()
	cachePath := "game:" + id
	game := Game{
		ID:      id,
		Players: make([]Player, 6),
	}

	err := s.cache.Set(r.Context(), cachePath, &game, time.Hour*24*90)
	if err != nil {
		errorResponse(w, http.StatusInternalServerError, err)
		return
	}

	http.Redirect(w, r, "/game/"+id, http.StatusFound)
}

func (s *Server) buildGameKey() string {
	key := uuid.New()
	return key.String()
}
