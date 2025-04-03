package models

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type Game struct {
	ID                               int         `json:"id"`
	Slug                             string      `json:"slug"`
	Gold                             int         `json:"gold"`
	Page                             int         `json:"page"`
	Notes                            string      `json:"notes"`
	CreatedAt                        time.Time   `json:"created_at"`
	Characters                       []Character `json:"players"`
	FlagBouncerBounced               bool        `json:"flag-bouncer-bounced"`
	FlagCulpritIdentified            bool        `json:"flag-culprit-identified"`
	FlagSlipping                     bool        `json:"flag-slipping"`
	FlagContractTerminated           bool        `json:"flag-contract-terminated"`
	FlagPrizedPosession              bool        `json:"flag-prized-posession"`
	FlagLegWound                     bool        `json:"flag-leg-wound"`
	FlagSoulJudgement                bool        `json:"flag-soul-judgement"`
	FlagFamilialDispute              bool        `json:"flag-familial-dispute"`
	FlagExperiencedGuide             bool        `json:"flag-experienced-guide"`
	FlagALonelySurvivor              bool        `json:"flag-a-lonely-survivor"`
	FlagMarked                       bool        `json:"flag-marked"`
	FlagHappyFugitives               bool        `json:"flag-happy-fugitives"`
	FlagHeartbrokenExile             bool        `json:"flag-heartbroken-exile"`
	FlagJudgementOfJudas             bool        `json:"flag-judgement-of-judas"`
	FlagThePatelDaughter             bool        `json:"flag-the-patel-daughter"`
	FlagEliasGratitude               bool        `json:"flag-elias-gratitude"`
	FlagAnEndToMadness               bool        `json:"flag-an-end-to-madness"`
	FlagNewAmbitions                 bool        `json:"flag-new-ambitions"`
	FlagAvoidingCollapse             bool        `json:"flag-avoiding-collapse"`
	FlagAHusbandsDuty                bool        `json:"flag-a-husbands-duty"`
	FlagASonsLove                    bool        `json:"flag-a-sons-love"`
	FlagKeyDefeat                    bool        `json:"flag-key-defeat"`
	FlagRested                       bool        `json:"flag-rested"`
	FlagBookOfGehenna                bool        `json:"flag-book-of-gehenna"`
	FlagFallenFriend                 bool        `json:"flag-fallen-friend"`
	FlagLastOfHerCoterie             bool        `json:"flag-last-of-her-coterie"`
	FlagSeedsOfHerPeople             bool        `json:"flag-seeds-of-her-people"`
	FlagGrievousWounds               bool        `json:"flag-grievous-wounds"`
	FlagTheOnlyWay                   bool        `json:"flag-the-only-way"`
	FlagIrreconcilableDifferences    bool        `json:"flag-irreconcilable-differences"`
	FlagReluctantPartnership         bool        `json:"flag-reluctant-partnership"`
	FlagMassacred                    bool        `json:"flag-massacred"`
	FlagCoffinBuddy                  bool        `json:"flag-coffin-buddy"`
	FlagCapableEnough                bool        `json:"flag-capable-enough"`
	FlagGoodRiddance                 bool        `json:"flag-good-riddance"`
	FlagTimeIsOfTheEssence           bool        `json:"flag-time-is-of-the-essence"`
	FlagAllWeCanGet                  bool        `json:"flag-all-we-can-get"`
	FlagButchered                    bool        `json:"flag-butchered"`
	FlagSacrifice                    bool        `json:"flag-sacrifice"`
	FlagAHelpingHand                 bool        `json:"flag-a-helping-hand"`
	FlagANewOwner                    bool        `json:"flag-a-new-owner"`
	FlagAGreaterCause                bool        `json:"flag-a-greater-cause"`
	FlagImprisonedAnointed           bool        `json:"flag-imprisoned-anointed"`
	AchievementNotQuiteUnkillable    bool        `json:"achievement-not-quite-unkillable"`
	AchievementDeathFromAbove        bool        `json:"achievement-death-from-above"`
	AchievementProofOfProwess        bool        `json:"achievement-proof-of-prowess"`
	AchievementLikeABlur             bool        `json:"achievement-like-a-blur"`
	AchievementSupremeMobility       bool        `json:"achievement-supreme-mobility"`
	AchievementSickleSlaughter       bool        `json:"achievement-sickle-slaughter"`
	AchievementRainingLoot           bool        `json:"achievement-raining-loot"`
	AchievementHealthyWealthyAndWise bool        `json:"achievement-healthy-wealthy-and-wise"`
	AchievementSuccessBreedsSuccess  bool        `json:"achievement-success-breeds-success"`
	AchievementZapAGap               bool        `json:"achievement-zap-a-gap"`
	AchievementPoppetBeGone          bool        `json:"achievement-poppet-be-gone"`
	AchievementToughWurm             bool        `json:"achievement-tough-wurm"`
	AchievementTentacleCollector     bool        `json:"achievement-tentacle-collector"`
	AchievementTorturedMortal        bool        `json:"achievement-tortured-mortal"`
	UpgradeMasterwork                bool        `json:"upgrade-masterwork"`
	UpgradeImbued                    bool        `json:"upgrade-imbued"`
	UpgradeDevastating               bool        `json:"upgrade-devastating"`
	UpgradeOtherworldly              bool        `json:"upgrade-otherworldly"`
	UpgradeReinforced                bool        `json:"upgrade-reinforced"`
	UpgradeBulky                     bool        `json:"upgrade-bulky"`
	UpgradeElegant                   bool        `json:"upgrade-elegant"`
	UpgradeRefined                   bool        `json:"upgrade-refined"`
	UpgradeEtherium                  bool        `json:"upgrade-etherium"`
	UpgradeSentient                  bool        `json:"upgrade-sentient"`
	UpgradeScrying                   bool        `json:"upgrade-scrying"`
	UpgradeDark                      bool        `json:"upgrade-dark"`
	UpgradeGlowing                   bool        `json:"upgrade-glowing"`
	UpgradeShimmering                bool        `json:"upgrade-shimmering"`
	UpgradeEnchanted                 bool        `json:"upgrade-enchanted"`
}

func GetGameBySlug(ctx context.Context, slug string) (Game, error) {
	ret := Game{}
	err := db.QueryRowContext(ctx, `
SELECT
	id,
	slug,
	gold,
	page,
	notes,
	created_at
FROM game
WHERE slug = $1
`, slug).Scan(
		&ret.ID,
		&ret.Slug,
		&ret.Gold,
		&ret.Page,
		&ret.Notes,
		&ret.CreatedAt,
	)

	if err != nil {
		return ret, fmt.Errorf("could not get game %s: %w", slug, err)
	}

	ret.Characters, err = GetCharactersForGame(ctx, ret.ID)
	if err != nil {
		return ret, fmt.Errorf("could not get characters for game %s: %w", slug, err)
	}

	return ret, nil
}

func (g *Game) Validate() error {
	var err error
	for _, char := range g.Characters {
		errors.Join(err, char.Validate())
	}

	return err
}

func (g *Game) Save(ctx context.Context) error {
	tx, err := db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	// First save the game
	if g.ID == 0 {
		err := g.insert(ctx, tx)
		if err != nil {
			return errors.Join(tx.Rollback(), fmt.Errorf("failed to insert game %q: %w", g.Slug, err))
		}
	} else {
		err := g.update(ctx, tx)
		if err != nil {
			return errors.Join(tx.Rollback(), fmt.Errorf("failed to update game %q: %w", g.Slug, err))
		}
	}

	// Next the characters
	for _, char := range g.Characters {
		char.GameID = g.ID
		if err := char.save(ctx, tx); err != nil {
			return errors.Join(tx.Rollback(), fmt.Errorf("failed to save game characters: %w", err))
		}
	}

	return tx.Commit()
}

func (g *Game) insert(ctx context.Context, tx *sql.Tx) error {
	stmt, err := tx.PrepareContext(ctx, `
INSERT INTO game (slug, gold, page, notes, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, NOW())
RETURNING id
`)

	if err != nil {
		return err
	}

	g.CreatedAt = time.Now().UTC()
	return stmt.QueryRowContext(ctx,
		&g.Slug,
		&g.Gold,
		&g.Page,
		&g.Notes,
		&g.CreatedAt,
	).Scan(&g.ID)
}

func (g *Game) update(ctx context.Context, tx *sql.Tx) error {
	stmt, err := tx.PrepareContext(ctx, `
UPDATE game SET
	gold = $2,
	page = $3,
	notes = $4,
	updated_at = NOW()
WHERE id = $1
RETURNING id
		`)

	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx,
		&g.ID,
		&g.Gold,
		&g.Page,
		&g.Notes,
	)
	return err
}
