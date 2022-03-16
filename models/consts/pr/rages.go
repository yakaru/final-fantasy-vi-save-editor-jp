package pr

import (
	"ffvi_editor/models/consts"
)

const (
	RageFrom int64 = 800
	RageTo   int64 = 1055
)

var (
	Rages = []*consts.NameValueChecked{
		consts.NewValueNameChecked(800, "Guard"),
		consts.NewValueNameChecked(801, "Imperial Soldier"),
		consts.NewValueNameChecked(802, "Templar"),
		consts.NewValueNameChecked(803, "Ninja"),
		consts.NewValueNameChecked(804, "Samurai"),
		consts.NewValueNameChecked(805, "Borghese"),
		consts.NewValueNameChecked(806, "Magna Roader"),
		consts.NewValueNameChecked(807, "Yojimbo"),
		consts.NewValueNameChecked(808, "Cloud"),
		consts.NewValueNameChecked(809, "Misty"),
		consts.NewValueNameChecked(810, "Al Jabr"),
		consts.NewValueNameChecked(811, "Zaghrem"),
		consts.NewValueNameChecked(812, "Apocrypha"),
		consts.NewValueNameChecked(813, "Dark Force"),
		consts.NewValueNameChecked(814, "Angel Whisper"),
		consts.NewValueNameChecked(815, "Oversoul"),
		consts.NewValueNameChecked(816, "Skeletal Horror"),
		consts.NewValueNameChecked(817, "Commander"),
		consts.NewValueNameChecked(818, "Mu"),
		consts.NewValueNameChecked(819, "Wererat"),
		consts.NewValueNameChecked(820, "Mugbear"),
		consts.NewValueNameChecked(821, "Belmodar"),
		consts.NewValueNameChecked(822, "Muud Suud"),
		consts.NewValueNameChecked(823, "Leaf Bunny"),
		consts.NewValueNameChecked(824, "Stray Cat"),
		consts.NewValueNameChecked(825, "Silver Lobo"),
		consts.NewValueNameChecked(826, "Doberman"),
		consts.NewValueNameChecked(827, "Megalodoth"),
		consts.NewValueNameChecked(828, "Fidor"),
		consts.NewValueNameChecked(829, "Briareus"),
		consts.NewValueNameChecked(830, "Suriander"),
		consts.NewValueNameChecked(831, "Chimera"),
		consts.NewValueNameChecked(832, "Behemoth"),
		consts.NewValueNameChecked(833, "Fafnir"),
		consts.NewValueNameChecked(834, "Lesser Lopros"),
		consts.NewValueNameChecked(835, "Fossil Dragon"),
		consts.NewValueNameChecked(836, "Holy Dragon"),
		consts.NewValueNameChecked(837, "Fiend Dragon"),
		consts.NewValueNameChecked(838, "Brachiosaur"),
		consts.NewValueNameChecked(839, "Tyrannosaur"),
		consts.NewValueNameChecked(840, "Darkwind"),
		consts.NewValueNameChecked(841, "Aepyornis"),
		consts.NewValueNameChecked(842, "Vulture"),
		consts.NewValueNameChecked(843, "Vasegiatta"),
		consts.NewValueNameChecked(844, "Zokka"),
		consts.NewValueNameChecked(845, "Trapper"),
		consts.NewValueNameChecked(846, "Hornet"),
		consts.NewValueNameChecked(847, "Nettlehopper"),
		consts.NewValueNameChecked(848, "Delta Beetle"),
		consts.NewValueNameChecked(849, "Killer Mantis"),
		consts.NewValueNameChecked(850, "Trillium"),
		consts.NewValueNameChecked(851, "Rafflesia"),
		consts.NewValueNameChecked(852, "Tumbleweed"),
		consts.NewValueNameChecked(853, "Vampire Thorn"),
		consts.NewValueNameChecked(854, "Cartagra"),
		consts.NewValueNameChecked(856, "Nautiloid"),
		consts.NewValueNameChecked(857, "Exocite"),
		consts.NewValueNameChecked(858, "Anguiform"),
		consts.NewValueNameChecked(859, "Leap Frog"),
		consts.NewValueNameChecked(860, "Lizard"),
		consts.NewValueNameChecked(861, "Litwor Chicken"),
		consts.NewValueNameChecked(862, "Slagworm"),
		consts.NewValueNameChecked(863, "Hell's Rider"),
		consts.NewValueNameChecked(865, "Onion Knight"),
		consts.NewValueNameChecked(866, "Magitek Armor"),
		consts.NewValueNameChecked(867, "Sky Armor"),
		consts.NewValueNameChecked(868, "Satellite"),
		consts.NewValueNameChecked(869, "Armored Weapon"),
		consts.NewValueNameChecked(870, "Spritzer"),
		consts.NewValueNameChecked(871, "Flan"),
		consts.NewValueNameChecked(872, "Outcast"),
		consts.NewValueNameChecked(873, "Humpty"),
		consts.NewValueNameChecked(874, "Brainpan"),
		consts.NewValueNameChecked(875, "Cruller"),
		consts.NewValueNameChecked(876, "Cactuar"),
		consts.NewValueNameChecked(877, "Bandit"),
		consts.NewValueNameChecked(878, "Harvester"),
		consts.NewValueNameChecked(879, "Bomb"),
		consts.NewValueNameChecked(880, "Still Life"),
		consts.NewValueNameChecked(881, "Lunatys"),
		consts.NewValueNameChecked(882, "Veil Dancer"),
		consts.NewValueNameChecked(883, "Hill Gigas"),
		consts.NewValueNameChecked(884, "Tonberry"),
		consts.NewValueNameChecked(885, "Magic Urn"),
		consts.NewValueNameChecked(886, "Mover"),
		consts.NewValueNameChecked(887, "Figaro Lizard"),
		consts.NewValueNameChecked(888, "Devoahan"),
		consts.NewValueNameChecked(889, "Aspiran"),
		consts.NewValueNameChecked(890, "Ghost"),
		consts.NewValueNameChecked(891, "Crawler"),
		consts.NewValueNameChecked(892, "Sand Ray"),
		consts.NewValueNameChecked(893, "Alacran"),
		consts.NewValueNameChecked(894, "Actinian"),
		consts.NewValueNameChecked(895, "Sandhorse"),
		consts.NewValueNameChecked(896, "Darkside"),
		consts.NewValueNameChecked(897, "Malboro"),
		consts.NewValueNameChecked(898, "Urok"),
		consts.NewValueNameChecked(899, "Foper"),
		consts.NewValueNameChecked(900, "Guard Leader"),
		consts.NewValueNameChecked(901, "Corporal"),
		consts.NewValueNameChecked(902, "General"),
		consts.NewValueNameChecked(903, "Covert"),
		consts.NewValueNameChecked(904, "Kamui"),
		consts.NewValueNameChecked(905, "Warlock"),
		consts.NewValueNameChecked(906, "Cherry"),
		consts.NewValueNameChecked(907, "Joker"),
		consts.NewValueNameChecked(908, "Iron Fist"),
		consts.NewValueNameChecked(909, "Devil"),
		consts.NewValueNameChecked(910, "Provoker"),
		consts.NewValueNameChecked(911, "Cloudwraith"),
		consts.NewValueNameChecked(912, "Mahadeva"),
		consts.NewValueNameChecked(913, "Vector Hound"),
		consts.NewValueNameChecked(914, "Peeper"),
		consts.NewValueNameChecked(915, "Stunner"),
		consts.NewValueNameChecked(916, "Sorath"),
		consts.NewValueNameChecked(917, "Destroyer"),
		consts.NewValueNameChecked(918, "Chippirabbit"),
		consts.NewValueNameChecked(919, "Coeurl Cat"),
		consts.NewValueNameChecked(920, "Bloodfang"),
		consts.NewValueNameChecked(921, "Hunting Hound"),
		consts.NewValueNameChecked(922, "Gorgias"),
		consts.NewValueNameChecked(923, "Don"),
		consts.NewValueNameChecked(924, "Murussu"),
		consts.NewValueNameChecked(925, "Wartpuck"),
		consts.NewValueNameChecked(926, "Gorgimera"),
		consts.NewValueNameChecked(927, "Behemoth King"),
		consts.NewValueNameChecked(928, "Vector Lythos"),
		consts.NewValueNameChecked(929, "Wyvern"),
		consts.NewValueNameChecked(930, "Zombie Dragon"),
		consts.NewValueNameChecked(931, "Dragon"),
		consts.NewValueNameChecked(932, "Primeval Dragon"),
		consts.NewValueNameChecked(933, "Weredragon"),
		consts.NewValueNameChecked(934, "Cirpius"),
		consts.NewValueNameChecked(935, "Sprinter"),
		consts.NewValueNameChecked(936, "Lenergia"),
		consts.NewValueNameChecked(937, "Marchosias"),
		consts.NewValueNameChecked(938, "Gloomwind"),
		consts.NewValueNameChecked(939, "Dropper"),
		consts.NewValueNameChecked(940, "Rock Wasp"),
		consts.NewValueNameChecked(941, "Grasswyrm"),
		consts.NewValueNameChecked(942, "Luridan"),
		consts.NewValueNameChecked(943, "Twinscythe"),
		consts.NewValueNameChecked(944, "Paraladia"),
		consts.NewValueNameChecked(945, "Exoray"),
		consts.NewValueNameChecked(946, "Crusher"),
		consts.NewValueNameChecked(947, "Ouroboros"),
		consts.NewValueNameChecked(948, "Acrophies"),
		consts.NewValueNameChecked(949, "Schmidt"),
		consts.NewValueNameChecked(950, "Devourer"),
		consts.NewValueNameChecked(951, "Cancer"),
		consts.NewValueNameChecked(952, "Gigantoad"),
		consts.NewValueNameChecked(953, "Basilisk"),
		consts.NewValueNameChecked(954, "Medusa Chicken"),
		consts.NewValueNameChecked(955, "Landworm"),
		consts.NewValueNameChecked(956, "Test Rider"),
		consts.NewValueNameChecked(957, "Pluto Armor"),
		consts.NewValueNameChecked(958, "Onion Dasher"),
		consts.NewValueNameChecked(959, "Heavy Armor"),
		consts.NewValueNameChecked(960, "Chaser"),
		consts.NewValueNameChecked(961, "Gamma"),
		consts.NewValueNameChecked(962, "Poplium"),
		consts.NewValueNameChecked(963, "Intangir"),
		consts.NewValueNameChecked(964, "Misfit"),
		consts.NewValueNameChecked(965, "Creature"),
		consts.NewValueNameChecked(966, "Enuo"),
		consts.NewValueNameChecked(967, "Deepeye"),
		consts.NewValueNameChecked(968, "Unseelie"),
		consts.NewValueNameChecked(969, "Neck Hunter"),
		consts.NewValueNameChecked(970, "Grenade"),
		consts.NewValueNameChecked(971, "Alluring Rider"),
		consts.NewValueNameChecked(972, "Pandora"),
		consts.NewValueNameChecked(973, "Blade Dancer"),
		consts.NewValueNameChecked(974, "Gigantos"),
		consts.NewValueNameChecked(975, "Magna Roader"),
		consts.NewValueNameChecked(976, "Lycaon"),
		consts.NewValueNameChecked(977, "Parasite"),
		consts.NewValueNameChecked(978, "Land Ray"),
		consts.NewValueNameChecked(979, "Antares"),
		consts.NewValueNameChecked(980, "Anemone"),
		consts.NewValueNameChecked(981, "Moonform"),
		consts.NewValueNameChecked(982, "Specter"),
		consts.NewValueNameChecked(983, "Great Malboro"),
		consts.NewValueNameChecked(984, "Bonnacon"),
		consts.NewValueNameChecked(985, "Oceanus"),
		consts.NewValueNameChecked(986, "Living Dead"),
		consts.NewValueNameChecked(988, "Face"),
		consts.NewValueNameChecked(989, "Outsider"),
		consts.NewValueNameChecked(990, "Coco"),
		consts.NewValueNameChecked(991, "Zeveak"),
		consts.NewValueNameChecked(992, "Nightwalker"),
		consts.NewValueNameChecked(993, "Demon Knight"),
		consts.NewValueNameChecked(994, "Imperial Elite"),
		consts.NewValueNameChecked(995, "Desert Hare"),
		consts.NewValueNameChecked(996, "Wizard"),
		consts.NewValueNameChecked(997, "Devil Fist"),
		consts.NewValueNameChecked(998, "Illuyankas"),
		consts.NewValueNameChecked(999, "Sergeant"),
		consts.NewValueNameChecked(1000, "Aspidochelon"),
		consts.NewValueNameChecked(1001, "Knotty"),
		consts.NewValueNameChecked(1002, "Luna Wolf"),
		consts.NewValueNameChecked(1003, "Belzecue"),
		consts.NewValueNameChecked(1004, "Caladrius"),
		consts.NewValueNameChecked(1005, "Tzakmaqiel"),
		consts.NewValueNameChecked(1006, "Lukhavi"),
		consts.NewValueNameChecked(1007, "Eukaryote"),
		consts.NewValueNameChecked(1008, "Land Grillon"),
		consts.NewValueNameChecked(1009, "Goetia"),
		consts.NewValueNameChecked(1010, "Greater Mantis"),
		consts.NewValueNameChecked(1011, "Bogy"),
		consts.NewValueNameChecked(1012, "Purusa"),
		consts.NewValueNameChecked(1013, "Black Dragon"),
		consts.NewValueNameChecked(1014, "Adamankary"),
		consts.NewValueNameChecked(1015, "Dante"),
		consts.NewValueNameChecked(1016, "Platinum Dragon"),
		consts.NewValueNameChecked(1017, "Duel Armor"),
		consts.NewValueNameChecked(1018, "Psychos"),
		consts.NewValueNameChecked(1019, "Mousse"),
		consts.NewValueNameChecked(1020, "Shambling Corpse"),
		consts.NewValueNameChecked(1021, "Punisher"),
		consts.NewValueNameChecked(1022, "Balloon"),
		consts.NewValueNameChecked(1023, "Gobbledygook"),
		consts.NewValueNameChecked(1024, "Great Behemoth"),
		consts.NewValueNameChecked(1025, "Scorpion"),
		consts.NewValueNameChecked(1026, "Chaos Dragon"),
		consts.NewValueNameChecked(1027, "Spitfire"),
		consts.NewValueNameChecked(1028, "Vector Chimera"),
		consts.NewValueNameChecked(1029, "Lich"),
		consts.NewValueNameChecked(1030, "Rukh"),
		consts.NewValueNameChecked(1031, "Magna Roader"),
		consts.NewValueNameChecked(1032, "Bug"),
		consts.NewValueNameChecked(1033, "Seaflower"),
		consts.NewValueNameChecked(1034, "Fortis"),
		consts.NewValueNameChecked(1035, "Venobennu"),
		consts.NewValueNameChecked(1036, "Galypdes"),
		consts.NewValueNameChecked(1037, "Junk"),
		consts.NewValueNameChecked(1038, "Mandrake"),
		consts.NewValueNameChecked(1039, "Valeor"),
		consts.NewValueNameChecked(1040, "Amduscias"),
		consts.NewValueNameChecked(1041, "Necromancer"),
		consts.NewValueNameChecked(1042, "Glasya Labolas"),
		consts.NewValueNameChecked(1043, "Magna Roader"),
		consts.NewValueNameChecked(1044, "Wild Rat"),
		consts.NewValueNameChecked(1045, "Gold Bear"),
		consts.NewValueNameChecked(1046, "InnoSent"),
		consts.NewValueNameChecked(1047, "Clymenus"),
		consts.NewValueNameChecked(1048, "Garm"),
		consts.NewValueNameChecked(1049, "Daedalus"),
		consts.NewValueNameChecked(1050, "Baalzephon"),
		consts.NewValueNameChecked(1051, "Ahriman"),
		consts.NewValueNameChecked(1052, "Death Machine"),
		consts.NewValueNameChecked(1053, "Metal Hitman"),
		consts.NewValueNameChecked(1054, "Io"),
		consts.NewValueNameChecked(1055, "Skip"),
	}
	SortedRages    []*consts.NameValueChecked
	RageLookupByID = make(map[int]*consts.NameValueChecked)
)

func init() {
	SortedRages = consts.SortByNameChecked(Rages)
	for _, r := range Rages {
		RageLookupByID[r.Value] = r
	}
}
