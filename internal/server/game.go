package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/taiidani/middara/internal/models"
)

type gameBag struct {
	Game models.Game
}

func (s *Server) gameHandler(w http.ResponseWriter, r *http.Request) {
	bag := gameBag{}

	slug := r.PathValue("slug")
	if len(slug) == 0 {
		errorResponse(w, http.StatusBadRequest, errInvalidGame)
		return
	}

	g, err := models.GetGameBySlug(r.Context(), slug)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	bag.Game = g

	// Ensure there are always 6 character slots, even if empty
	for i := len(bag.Game.Characters) - 1; i < 5; i++ {
		bag.Game.Characters = append(bag.Game.Characters, models.Character{})
	}

	template := "game.gohtml"
	renderHtml(w, http.StatusOK, template, bag)
}

func (s *Server) saveGameHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id == 0 {
		errorResponse(w, http.StatusBadRequest, errGameIDRequired)
		return
	}

	gold, err := strconv.Atoi(r.FormValue("gold"))
	if err != nil {
		errorResponse(w, http.StatusBadRequest, errInvalidGame)
		return
	}

	page, err := strconv.Atoi(r.FormValue("page"))
	if err != nil {
		errorResponse(w, http.StatusBadRequest, errInvalidGame)
		return
	}

	game := models.Game{
		ID:                               id,
		Slug:                             r.FormValue("slug"),
		Gold:                             gold,
		Page:                             page,
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

	for i := range 6 {
		if r.FormValue(fmt.Sprintf("name-%d", i)) == "" {
			continue
		}

		id, _ := strconv.Atoi(r.FormValue(fmt.Sprintf("id-%d", i)))

		xp, err := strconv.Atoi(r.FormValue(fmt.Sprintf("xp-%d", i)))
		if err != nil {
			errorResponse(w, http.StatusBadRequest, errInvalidCharacter)
			return
		}

		damage, err := strconv.Atoi(r.FormValue(fmt.Sprintf("damage-%d", i)))
		if err != nil {
			errorResponse(w, http.StatusBadRequest, errInvalidCharacter)
			return
		}

		game.Characters = append(game.Characters, models.Character{
			ID:           id,
			Name:         r.FormValue(fmt.Sprintf("name-%d", i)),
			XP:           xp,
			Damage:       damage,
			Injured:      r.FormValue(fmt.Sprintf("injured-%d", i)) == "on",
			Unselectable: r.FormValue(fmt.Sprintf("unselectable-%d", i)) == "on",
		})
	}

	err = game.Validate()
	if err != nil {
		errorResponse(w, http.StatusBadRequest, err)
		return
	}

	err = game.Save(r.Context())
	if err != nil {
		errorResponse(w, http.StatusInternalServerError, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/game/%s", game.Slug), http.StatusFound)
}

func (s *Server) newGameHandler(w http.ResponseWriter, r *http.Request) {
	slug := s.buildGameKey()
	game := models.Game{
		Slug:       slug,
		Characters: []models.Character{},
	}

	err := game.Validate()
	if err != nil {
		errorResponse(w, http.StatusBadRequest, err)
		return
	}

	err = game.Save(r.Context())
	if err != nil {
		errorResponse(w, http.StatusInternalServerError, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/game/%s", game.Slug), http.StatusFound)
}

func (s *Server) buildGameKey() string {
	key := uuid.New()
	return key.String()
}
